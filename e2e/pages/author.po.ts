import { browser, by, element, promise } from 'protractor';

export class AuthorPage {
  navigateTo() {
    return browser.get('/author');
  }
}
