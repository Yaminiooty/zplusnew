pipeline {
    agent any

    stages {
        stage('Checkout and Run Docker') {
            steps {
                script {
                    checkout([$class: 'GitSCM', branches: [[name: 'develop']], userRemoteConfigs: [[url: 'https://github.com/Yaminiooty/zplusnew.git']]])
                    sh '''
                        cd zplusnew
                        pwd
                        docker-compose up -d
                    '''
                }
            }
        }
    }
}
