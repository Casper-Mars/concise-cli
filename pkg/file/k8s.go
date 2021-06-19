package file

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
)

func BuildK8sFile(path string) error {
	group, _ := errgroup.WithContext(context.Background())
	path = path + "/k8s"
	group.Go(func() error {
		return initDPFile(path)
	})
	group.Go(func() error {
		return initServiceFile(path)
	})
	group.Go(func() error {
		return initIngressFile(path)
	})
	return group.Wait()
}

func initDPFile(path string) error {
	target, err := os.OpenFile(path+"/dp.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: ${appName}-backend
  namespace: ${namespace}
  labels:
    k8s-app: ${appName}-backend
spec:
  replicas: 2
  selector:
    matchLabels:
      k8s-app: ${appName}-backend
  template:
    metadata:
      labels:
        version: VERSION_PLACEHOLDER
        k8s-app: ${appName}-backend
      annotations:
        prometheus.io/port: http-metrics
        prometheus.io/scrape: "true"
        prometheus.io/path: /actuator/prometheus
    spec:
      nodeSelector:
        role: worker
      tolerations:
        - key: "worker"
          operator: "Exists"
          effect: "NoSchedule"
      containers:
        - name: ${appName}-backend
          image: IMAGE_PLACEHOLDER
          ports:
            - containerPort: 8080
              protocol: TCP
          livenessProbe:
            initialDelaySeconds: 5
            periodSeconds: 5
            exec:
              command:
                - ls
          readinessProbe:
            httpGet:
              path: /test/ping
              port: 8080`)
	return err
}

func initServiceFile(path string) error {
	target, err := os.OpenFile(path+"/svc.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
apiVersion: v1
kind: Service
metadata:
  name: ${appName}-backend
  namespace: ${namespace}
  labels:
    k8s-app: ${appName}-backend
spec:
  selector:
    k8s-app: ${appName}-backend
  ports:
    - port: 80
      targetPort: 8080`)
	return err
}

func initIngressFile(path string) error {
	target, err := os.OpenFile(path+"/ingress.yaml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: ${appName}-backend
  namespace: ${namespace}
spec:
  entryPoints:
    - web
  routes:
    - match: Host() && PathPrefix()
      kind: Rule
      middlewares:
        - name: ${appName}-stripprefix
      services:
        - name: ${appName}-backend
          port: 80
---
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: ${appName}-stripprefix
  namespace: ${namespace}
spec:
  stripPrefix:
    prefixes:
      - /api`)
	return err
}
