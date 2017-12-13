import { browser, by, element, promise, ElementFinder } from 'protractor';
import * as utils from '../utils/utils';

export interface UserReg {
  Email: string;
  Password: string;
  UserName: string;
}

export interface UserLogin {
  Email: string;
  Password: string;
}

export class LoginPage {
  register(user: UserReg) {
    utils.type('#register-email', user.Email);
    utils.type('#register-username', user.UserName);
    utils.type('#register-password', user.Password);
    utils.type('#register-password-confirmation', user.Password);
    utils.e('#register-submit').click();
  }

  login(user: UserLogin) {
    utils.type('#login-email', user.Email);
    utils.type('#login-password', user.Password);
    utils.e('#login-submit').click();
  }
}
