pipeline {
    agent any
    tools {
        go 'go-lang' 
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go1.git'
                
            }
        }
        
        stage('Test') {
            steps {
                sh 'go test'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o prog.sh prog.go'
            }
        }

        stage('Deploy') {
            steps {
                sh'scp -i key .prog.sh loco2@192.168.81.130:/home/loco2' 
                
                    
           
               
            }
        }
    }
}
