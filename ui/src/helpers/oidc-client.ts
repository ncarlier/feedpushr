import { Log, UserManager, UserManagerSettings } from 'oidc-client'

export class OIDCClient {
  public userManager: UserManager
  private origin: string

  constructor(settings: Partial<UserManagerSettings>, origin = document.location.origin) {
    settings = {
      response_type: 'code',
      scope: 'openid profile email',
      ...settings,
    }
    this.origin = origin
    this.userManager = new UserManager(settings)
    this.userManager.events.addUserSignedOut(async () => {
      console.log('user signed out...')
      this.logout()
    })
    this.userManager.clearStaleState()

    Log.logger = console
    Log.level = Log.WARN
  }

  public getUser() {
    return this.userManager.getUser()
  }

  public getAccountUrl() {
    const { authority, client_id } = this.userManager.settings
    return `${authority}/account?referrer=${client_id}&referrer_uri=${encodeURI(this.origin)}`
  }

  public login() {
    return this.userManager.signinRedirect()
  }

  public renewToken() {
    return this.userManager.signinSilent()
  }

  public logout() {
    return this.userManager.signoutRedirect()
  }
}
