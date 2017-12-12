import { HomePage } from './pages/home.po';
import { LoginPage } from './pages/login.po';
import { Header } from './pages/header.po';
import {
  browser,
  by,
  element,
  promise,
  ElementFinder,
  protractor
} from 'protractor';

function makeid() {
  let text = '';
  const possible = 'abcdfghijklmnopqstuvwxyz012346789';
  for (let i = 0; i < 5; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

describe('Home page', () => {
  let homePage: HomePage;
  let loginPage: LoginPage;
  let header: Header;
  const id = makeid();

  browser.driver
    .manage()
    .window()
    .setSize(1200, 800);

  beforeAll(() => {
    browser.waitForAngularEnabled(false); // the demo page on the homepage is async and trips protractor up
    homePage = new HomePage();
    loginPage = new LoginPage();
    header = new Header();
    const Id = this.id;
  });

  it('should be able to register', () => {
    homePage.navigateTo();
    header.register({
      UserName: 'user-' + id,
      Email: 'user' + id + '@gmail.com',
      Password: 'pw'
    });
  });

  it('should be able to login', () => {
    header.logout().then(() => {
      header.login({
        Email: 'user' + id + '@gmail.com',
        Password: 'pw',
        UserName: 'user-' + id
      });
    });
  });
});
