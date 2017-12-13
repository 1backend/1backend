import { HomePage } from './pages/home.po';
import { LoginPage } from './pages/login.po';
import { Header } from './pages/header.po';
import { AuthorPage } from './pages/author.po';
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
  const possible = 'abcdefghijklmnopqrstuvwxyz0123456789';
  for (let i = 0; i < 12; i++) {
    text += possible.charAt(Math.floor(Math.random() * possible.length));
  }
  return text;
}

describe('Home page', () => {
  let homePage: HomePage;
  let loginPage: LoginPage;
  let header: Header;
  let authorPage: AuthorPage;
  const id = makeid();
  const pw = makeid();
  console.log('generated id:', id);

  browser.driver
    .manage()
    .window()
    .setSize(1200, 800);

  beforeAll(() => {
    browser.waitForAngularEnabled(false); // the demo page on the homepage is async and trips protractor up
    homePage = new HomePage();
    loginPage = new LoginPage();
    header = new Header();
    authorPage = new AuthorPage();
  });

  it('should be able to register', () => {
    homePage.navigateTo();
    header.register({
      UserName: 'user-' + id,
      Email: 'user' + id + '@gmail.com',
      Password: pw
    });
  });

  // uncomment this to test it locally - can't test payment in live
  // it('should be able to top up', () => {
  //  authorPage.topUpBy(9);
  // });

  it('should be able to login', () => {
    header.logout().then(() => {
      header.login({
        Email: 'user' + id + '@gmail.com',
        Password: pw,
        UserName: 'user-' + id
      });
    });
  });
});
