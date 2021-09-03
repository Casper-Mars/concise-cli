package template

//NewGitIgnore return .gitignore file template
func NewGitIgnore() string {
	return `# Should be edit by user
target/
.settings/
.idea/
.project
.classpath
.springBeans
*.iml
./hack-temp
.DS_Store
`
}
