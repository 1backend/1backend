import { Component, OnInit, ViewChild, ElementRef, Renderer, AfterViewInit } from '@angular/core';

@Component({
  selector: 'app-demo',
  templateUrl: './demo.component.html',
  styleUrls: ['./demo.component.css']
})
export class DemoComponent implements OnInit, AfterViewInit {
  @ViewChild('saveButton') saveButton: ElementRef;

  title = 'Pick tech'
  page = 'language-select'
  selected = 0;
  code = `func hello(rsp http.Response) {
    write(rsp, "hi there")
}`

  terminal = `me@laptop$ curl 1backend.com/app/me/my-app/hello
hi there`
  counter = 0;

  constructor(
    private renderer: Renderer
  ) {}

  typeTerminal() {

  }

  ngOnInit() {

  }

  ngAfterViewInit() {
    this.animate(this);
    const that = this;
    const f = this.animate;
    setInterval(() => {
      f(that)
    }, 9000);
  }

  triggerFalseClick(that: DemoComponent) {
    // this crap does not work
    // let el: HTMLElement = that.saveButton.nativeElement as HTMLElement;
    // el.click();
  }

  reset() {
    this.counter = 0;
  }

  wait(n: number): number {
    this.counter += n;
    return this.counter;
  }

  animate(that: DemoComponent) {
    that.reset();
    setTimeout(() => {
      that.selected = 2;
    }, that.wait(2000))
    setTimeout(() => {
      that.page = 'app';
      that.title = 'Customize your endpoints';
      that.selected = 0;
    }, that.wait(1000))
    setTimeout(() => {
      that.triggerFalseClick(that);
    }, that.wait(2000))
    setTimeout(() => {
      that.page = 'terminal';
      that.title = 'Your API is live!';
    }, that.wait(1000))
    setTimeout(() => {
      that.page = 'language-select';
      that.title = 'Pick tech';
    }, that.wait(3000))
  }
}
