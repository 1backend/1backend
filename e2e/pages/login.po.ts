import { browser, by, element, promise, ElementFinder } from 'protractor';

export interface UserReg {
  Email: string;
  Password: string;
  UserName: string;
}

function e(s: string): ElementFinder {
  return element(by.css(s));
}

function type(id: string, text: string) {
  const el = element(by.id(id));
  el.click();
  el.sendKeys(text);
}

export class LoginPage {
  register(user: UserReg) {
    type('register-email', user.Email);
    type('register-username', user.UserName);
    type('register-password', user.Password);
    type('register-password-confirmation', user.Password);
    e('#register-submit').click();
  }
}
