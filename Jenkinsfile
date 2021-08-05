pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        git(url: 'https://github.com/nengbai/study.git', branch: 'master')
      }
    }

    stage('test') {
      steps {
        sh 'ls -ltr'
      }
    }

  }
  environment {
    CI = 'True'
  }
}