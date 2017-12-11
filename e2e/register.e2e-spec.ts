import { HomePage } from './pages/home.po';
import { browser, by, element, promise, ElementFinder } from 'protractor';

function makeid() {
  let text = '';
  const possible =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  for (let i = 0; i < 5; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

describe('Register from home page', () => {
  let homePage: HomePage;
  const id = makeid();
  browser.driver
    .manage()
    .window()
    .maximize();

  beforeEach(() => {
    homePage = new HomePage();
  });

  it('should have register button', () => {
    homePage.navigateTo();
    expect(homePage.hasRegisterButton()).toBeTruthy();

    // homePage.clickRegister().then(() => {
    // const regTab = element(by.css('#register-tab'));
    //  expect(regTab.isPresent()).toBeTruthy();
    // });
  });
});
