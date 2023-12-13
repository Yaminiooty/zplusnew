pipeline {

   agent any
   
     stages {
       stage('checkout') {
           steps {
               sh '''
                sudo su
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
