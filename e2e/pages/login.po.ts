import { browser, by, element, promise, ElementFinder } from 'protractor';

export interface UserReg {
  Email: string;
  Password: string;
  UserName: string;
}

function e(s: string): ElementFinder {
  return element(by.css(s));
}

export class LoginPage {
  register(user: UserReg): promise.Promise<void> {
    return e('#register-email')
      .sendKeys(user.Email)
      .then(() => {
        e('#register-username').sendKeys(user.UserName);
      })
      .then(() => {
        e('#register-password').sendKeys(user.Password);
      })
      .then(() => {
        e('#register-password-confirmation').sendKeys(user.Password);
      })
      .then(() => {
        e('#register-submit').click();
      });
  }
}
