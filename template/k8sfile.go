package template

func NewK8sFile() string {
	return `# Should edit
#------------------------------svc-------------------------------
# Define service that export ports 80 and 9933
apiVersion: v1
kind: Service
metadata:
  name: #service name#
  namespace: #namespace#
  labels:
    k8s-app: #service labels#
spec:
  selector:
    k8s-app: #pod labels#
  ports:
    - name: http
      port: 80
      protocol: TCP
      targetPort: 8080
    - name: xxl
      port: 9933
      protocol: TCP
      targetPort: 9933

---
#------------------------------ingress-----------------------------
apiVersion: traefik.containo.us/v1alpha1
kind: IngressRoute
metadata:
  name: #ingress name#
  namespace: #namespace#
spec:
  entryPoints:
    - web
  routes:
    - match: Host() && PathPrefix()
      kind: Rule
      middlewares:
        - name: api-stripprefix
      services:
        - name: #service name#
          port: 80

---
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  name: api-stripprefix
  namespace: #namespace#
spec:
  stripPrefix:
    prefixes:
      - /api

---
#------------------------------dp---------------------------------
apiVersion: apps/v1
kind: Deployment
metadata:
  name: #dp name#
  namespace: #namespace#
  labels:
    k8s-app: #dp labels#
spec:
  replicas: 2
  selector:
    matchLabels:
      k8s-app: #pod labels#
  template:
    metadata:
      labels:
        version: VERSION_PLACEHOLDER
        k8s-app: #pod labels#
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
        - name: 
          image: 
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
              path: /heath/ping
              port: 8080
`
}
