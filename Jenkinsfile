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
                
                cd /home/rocky
                git clone https://github.com/Yaminiooty/zplusnew.git 
                cd zplusnew
                pwd
                docker-compose up -d
                          '''   
                  }
           }
       
   }
}  
