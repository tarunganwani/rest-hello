pipeline{
	agent any
	
// 	environment {
// 	     GOPATH ="""C:\\Users\\tarun.ganwani\\go"""
//     }
    
	stages{
		stage('Build'){
		    steps{
    		    bat "set GOPATH=C:\\Users\\tarun.ganwani\\go; cd basic-todo; go build"
    		    bat "set GOPATH=C:\\Users\\tarun.ganwani\\go; cd model-handler-app; go build"
    		    bat "set GOPATH=C:\\Users\\tarun.ganwani\\go; cd tdd-model-handler; go build"
		    }
		}
		stage('Test'){
		    steps{
    		    bat "set GOPATH=C:\\Users\\tarun.ganwani\\go; cd tdd-model-handler\\test; go test"
		    }
		}
		stage('Deploy'){
		    steps{
    		    bat "echo Deploying application"
		    }
		}
	}
}