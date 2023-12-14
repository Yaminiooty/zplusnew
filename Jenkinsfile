pipeline {

   agent any
   
     stages {
      stage('Check Jenkins User') {
          steps {
                 sh 'whoami'
                }
            }

       stage('checkout') {
           steps {
               sh '''
               
                cd /
                git clone -b develop https://github.com/Yaminiooty/zplusnew.git 
                sudo chmod +x -R /zplusnew
                sudo cd zplusnew
                pwd
                docker-compose up -d
                          '''   
                  }
           }
       
   }
}  
