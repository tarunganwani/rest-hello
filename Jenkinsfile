pipeline{
	agent any
	
	environment {
        GOPATH = "${pwd}"
        PATH = "C:\\Users\\tarun.ganwani\\AppData\\Local\\Programs\\Git\\usr\\bin;C:\\Users\\tarun.ganwani\\AppData\\Local\\Programs\\Git\\bin;${env.PATH}"
    }
    
	stages{
		stage('Build'){
		    steps{
    		    bat "cd basic-todo"
    			bat "go build"
		    }
		}
	}
}