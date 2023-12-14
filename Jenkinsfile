pipeline {
  agent any
  stages {
    stage('one') {
      steps {
        git(url: 'https://github.com/chaozai0304/golearning.git', branch: 'main')
        sh 'ls'
        sh '''ls -l
'''
      }
    }

    stage('two') {
      steps {
        sh 'uname -a'
      }
    }

  }
}