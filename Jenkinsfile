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
        sh './jenkins/scripts/test.sh'
      }
    }

  }
  environment {
    CI = 'True'
  }
}