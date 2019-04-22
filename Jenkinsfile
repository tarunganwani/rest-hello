pipeline{
	agent any
	
// 	environment {
// 	     GOPATH ="""C:\\Users\\tarun.ganwani\\go"""
//     }
    
	stages{
		stage('Build'){
		    steps{
    		    bat "set GOPATH=C:\\Users\\tarun.ganwani\\go; cd basic-todo; go build"
		    }
		}
	}
}