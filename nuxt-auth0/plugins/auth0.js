import Auth0Lock from 'auth0-lock'
import nuxtConfig from '~/nuxt.config'
const config = nuxtConfig.auth0

class Auth0Util {
  showLock(container) {
    const lock = new Auth0Lock(
      config.clientID,
      config.domain,
      {
        container,
        closable: false,
        auth: {
          responseType: 'token id_token',
          redirectUrl: this.getBaseUrl() + '/callback',
          params: {
            scope: 'openid profile email'
          }
        }
      }
    )

    lock.show()
  }

  getBaseUrl() {
    return `${window.location.protocol}//${window.location.host}`
  }
}

export default (context, inject) => {
  inject('auth0', new Auth0Util)
}
