pipeline{
	agent any
	
	stages{
		stage('Build'){
			sh 'set GOPATH=C:\Users\tarun.ganwani\go'
			sh 'cd basic-todo'
			sh 'go build' 

			sh 'cd model-handler-app'
			sh 'go build'
			
			sh 'cd tdd-model-handler'
			sh 'go build'
		}
	}
}