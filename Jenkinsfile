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
                sudo chmod +x -R /home/rocky/test
                cd /home/rocky/test
                git clone -b develop https://github.com/Yaminiooty/zplusnew.git 
                
                sudo cd zplusnew
                pwd
                docker-compose up -d
                          '''   
                  }
           }
       
   }
}  
