pipeline {
  agent none
  stages {
    stage('one') {
      parallel {
        stage('one') {
          steps {
            sh 'ls'
            sh '''ls -l
'''
          }
        }

        stage('') {
          steps {
            echo '123'
          }
        }

      }
    }

    stage('two') {
      steps {
        sh 'uname -a'
      }
    }

  }
}