pipeline {
    agent any
 

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go.git'
            }
        }

        
stage('Test') {
            steps {
                bat 'go test'
            }
        }

        stage('Build') {
            steps {
                bat 'go build -o prog.exe prog.go'
            }
        }

        stage('Run') {
            steps {
                bat '.\\prog.exe'
            }
        }
    }
}

