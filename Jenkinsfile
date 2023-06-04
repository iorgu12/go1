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
                bat 'go build -o prog.sh prog.go'
            }
        }

        stage('Deploy') {
            steps {
                script {
                    def remote = [:]
                    remote.name = 'my-ssh-server'
                    remote.host = '192.168.81.129'
                    remote.user = 'coco'
                    remote.password = 'loco12'  
                    remote.allowAnyHosts = true   
                    echo 'Deploying....'
                    sshPut remote: remote, from: 'prog.sh', into: '/home/coco/lab'
                    sshCommand remote: remote, command:"chmod +x /home/coco/lab/prog.sh"
                    sshCommand remote: remote, command:"systemctl start gop.service"
                    
           
                }
            }
        }
    }
}
