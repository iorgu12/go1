pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/iorgu12/go1.git'
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

        stage('Deploy') {
            steps {
                script {
                    withCredentials([sshUserPrivateKey(credentialsId: 'f3837549-e013-4e53-8f80-077e04b7b058', keyFileVariable: 'popo')]) {
                        echo 'Deploying....'
                        bat """
                            scp -i $popo prog.exe coco@192.168.81.129:/home/iorgu/lab
                            ssh -i $popo coco@192.168.81.129 "chmod +x /home/iorgu/lab/prog.exe && /home/iorgu/lab/prog.exe"
                        """
                    }
                }
            }
        }
    }
}
