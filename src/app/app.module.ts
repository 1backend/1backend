import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AceEditorModule } from 'ng2-ace-editor';
import 'brace/mode/sql';
// import 'brace/mode/ada';
// import 'brace/mode/apache_conf';
// import 'brace/mode/applescript';
// import 'brace/mode/asciidoc';
// import 'brace/mode/assembly_x86';
// import 'brace/mode/autohotkey';
// import 'brace/mode/batchfile';
// import 'brace/mode/bro';
// import 'brace/mode/c9search';
// import 'brace/mode/c_cpp';
// import 'brace/mode/cirru';
// import 'brace/mode/clojure';
// import 'brace/mode/cobol';
// import 'brace/mode/coffee';
// import 'brace/mode/coldfusion';
// import 'brace/mode/csharp';
// import 'brace/mode/css';
// import 'brace/mode/curly';
// import 'brace/mode/dart';
// import 'brace/mode/django';
// import 'brace/mode/d';
// import 'brace/mode/dockerfile';
// import 'brace/mode/dot';
// import 'brace/mode/drools';
// import 'brace/mode/eiffel';
// import 'brace/mode/ejs';
// import 'brace/mode/elixir';
// import 'brace/mode/elm';
// import 'brace/mode/erlang';
// import 'brace/mode/forth';
// import 'brace/mode/fortran';
// import 'brace/mode/ftl';
// import 'brace/mode/gcode';
// import 'brace/mode/gherkin';
// import 'brace/mode/gitignore';
// import 'brace/mode/glsl';
// import 'brace/mode/gobstones';
import 'brace/mode/golang';
// import 'brace/mode/groovy';
// import 'brace/mode/haml';
// import 'brace/mode/handlebars';
// import 'brace/mode/haskell_cabal';
// import 'brace/mode/haskell';
// import 'brace/mode/haxe';
// import 'brace/mode/hjson';
// import 'brace/mode/html_elixir';
// import 'brace/mode/html';
// import 'brace/mode/html_ruby';
// import 'brace/mode/ini';
// import 'brace/mode/io';
// import 'brace/mode/jack';
// import 'brace/mode/jade';
// import 'brace/mode/java';
import 'brace/mode/javascript';
// import 'brace/mode/jsoniq';
import 'brace/mode/json';
// import 'brace/mode/jsp';
// import 'brace/mode/jsx';
// import 'brace/mode/julia';
// import 'brace/mode/kotlin';
// import 'brace/mode/latex';
// import 'brace/mode/less';
// import 'brace/mode/liquid';
// import 'brace/mode/lisp';
// import 'brace/mode/logiql';
// import 'brace/mode/lsl';
// import 'brace/mode/lua';
// import 'brace/mode/luapage';
// import 'brace/mode/lucene';
// import 'brace/mode/makefile';
import 'brace/mode/markdown';
// import 'brace/mode/mask';
// import 'brace/mode/matlab';
// import 'brace/mode/maze';
// import 'brace/mode/mel';
// import 'brace/mode/mushcode';
// import 'brace/mode/mysql';
// import 'brace/mode/nix';
// import 'brace/mode/nsis';
// import 'brace/mode/objectivec';
// import 'brace/mode/ocaml';
// import 'brace/mode/pascal';
// import 'brace/mode/perl';
// import 'brace/mode/pgsql';
// import 'brace/mode/php';
// import 'brace/mode/powershell';
// import 'brace/mode/praat';
// import 'brace/mode/prolog';
// import 'brace/mode/properties';
// import 'brace/mode/protobuf';
// import 'brace/mode/python';
// import 'brace/mode/razor';
// import 'brace/mode/rdoc';
// import 'brace/mode/rhtml';
// import 'brace/mode/r';
// import 'brace/mode/rst';
// import 'brace/mode/ruby';
// import 'brace/mode/rust';
// import 'brace/mode/sass';
// import 'brace/mode/scad';
// import 'brace/mode/scala';
// import 'brace/mode/scheme';
// import 'brace/mode/scss';
import 'brace/mode/sh';
// import 'brace/mode/sjs';
// import 'brace/mode/smarty';
// import 'brace/mode/snippets';
// import 'brace/mode/soy_template';
// import 'brace/mode/space';
// import 'brace/mode/sql';
// import 'brace/mode/sqlserver';
// import 'brace/mode/stylus';
// import 'brace/mode/svg';
// import 'brace/mode/swift';
// import 'brace/mode/tcl';
// import 'brace/mode/tex';
// import 'brace/mode/textile';
// import 'brace/mode/toml';
// import 'brace/mode/tsx';
// import 'brace/mode/twig';
import 'brace/mode/typescript';
// import 'brace/mode/vala';
// import 'brace/mode/vbscript';
// import 'brace/mode/velocity';
// import 'brace/mode/verilog';
// import 'brace/mode/vhdl';
// import 'brace/mode/wollok';
// import 'brace/mode/xml';
// import 'brace/mode/xquery';
// import 'brace/mode/yaml';
// import 'brace/mode/abap';
// import 'brace/mode/abc';
// import 'brace/mode/actionscript';
// import 'brace/mode/lean';
// import 'brace/mode/live_script';
// import 'brace/mode/livescript';
// import 'brace/mode/mavens_mate_log';
// import 'brace/mode/mips_assembler';
// import 'brace/mode/mipsassembler';
// import 'brace/mode/swig';
// import 'brace/mode/diff';
// import 'brace/mode/plain_text';
// import 'brace/mode/text';
// import 'brace/index';
// import * as ace from 'brace';
// import 'brace/theme/github';
// import 'brace/theme/sqlserver';

import { AppComponent } from './app.component';
import { ProjectComponent } from './projects/project/project.component';
import { RouterModule, Routes } from '@angular/router';
import { AuthorComponent } from './author/author.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {
  MatTooltipModule,
  MatMenuModule,
  MatSnackBar,
  MatToolbarModule,
  MatPaginatorModule,
  MatProgressSpinnerModule,
  MatCardModule,
  MatFormFieldModule,
  MatInputModule,
  MatSelectModule,
  MatOptionModule,
  MatButtonModule,
  MatIconModule,
  MatTabsModule,
  MatDialogModule,
  MatDialog,
  MatListModule,
  MatExpansionModule,
  MatCheckboxModule,
} from '@angular/material';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { FlexLayoutModule } from '@angular/flex-layout';
import { LoginComponent } from './login/login.component';
import { SimpleNotificationsModule } from 'angular2-notifications';
import { HomeComponent } from './home/home.component';
import { MomentModule } from 'angular2-moment';
import { SqlComponent } from './projects/project/sql/sql.component';
import { BuildsComponent } from './projects/project/builds/builds.component';
import { CodeComponent } from './projects/project/code/code.component';
import { MarkdownToHtmlModule } from 'ng2-markdown-to-html';
import { SessionService } from './session.service';
import { ConstService } from './const.service';
import { ChargeService } from './charge.service';
import { DemoComponent } from './home/demo/demo.component';
import { HeaderComponent } from './header/header.component';
import { ProjectsComponent } from './projects/projects.component';
import { FilterPipe } from './filter.pipe';
import { LoginDialogService } from './login/login-dialog.service';
import { UserService } from './user.service';
import { CreateProjectComponent } from './projects/create-project/create-project.component';
import { Ng2PageScrollModule } from 'ng2-page-scroll';
import { CreateProjectDialogComponent } from './projects/create-project-dialog/create-project-dialog.component';
import { CreateProjectDialogService } from './projects/create-project-dialog.service';
import 'hammerjs';
import { IssueComponent } from './projects/project/issues/issue/issue.component';
import { CreateIssueDialogService } from './projects/project/issues/create-issue-dialog.service';
import { IssuesComponent } from './projects/project/issues/issues.component';
import { PricingComponent } from './home/pricing/pricing.component';
import { CreateIssueComponent } from './projects/project/issues/create-issue/create-issue.component';
import { ProjectListComponent } from './projects/project-list/project-list.component';
import { LengthPipe } from './length.pipe';
import { ResetComponent } from './reset/reset.component';
import { RecoverComponent } from './recover/recover.component';
import { PostsComponent } from './author/posts/posts.component';
import { TokensComponent } from './author/tokens/tokens.component';
import { ProfileEditComponent } from './author/profile-edit/profile-edit.component';
import { TokenService } from './token.service';
import { ProjectService } from './project.service';

const appRoutes: Routes = [
  {
    path: 'reset',
    component: ResetComponent
  },
  {
    path: 'recover',
    component: RecoverComponent
  },
  {
    path: 'projects',
    component: ProjectsComponent
  },
  {
    path: ':author/:project/issue/:issueId',
    component: ProjectComponent
  },
  {
    path: ':author/:project',
    component: ProjectComponent
  },
  {
    path: ':author/:project/:tab',
    component: ProjectComponent
  },
  {
    path: 'demo',
    component: DemoComponent
  },
  {
    path: ':author',
    component: AuthorComponent
  },
  {
    path: '',
    component: HomeComponent
  }
];

@NgModule({
  declarations: [
    AppComponent,
    ProjectComponent,
    AuthorComponent,
    LoginComponent,
    SqlComponent,
    BuildsComponent,
    CodeComponent,
    HomeComponent,
    DemoComponent,
    HeaderComponent,
    ProjectsComponent,
    FilterPipe,
    CreateProjectComponent,
    CreateProjectDialogComponent,
    IssuesComponent,
    IssueComponent,
    PricingComponent,
    CreateIssueComponent,
    ProjectListComponent,
    LengthPipe,
    ResetComponent,
    RecoverComponent,
    PostsComponent,
    TokensComponent,
    ProfileEditComponent,

  ],
  imports: [
    BrowserModule,
    MarkdownToHtmlModule.forRoot(),
    FormsModule,
    RouterModule.forRoot(
      appRoutes,
      { enableTracing: true } // <-- debugging purposes only
    ),
    AceEditorModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatOptionModule,
    MatButtonModule,
    MatIconModule,
    MatTabsModule,
    MatPaginatorModule,
    MatProgressSpinnerModule,
    MatToolbarModule,
    FlexLayoutModule,
    MomentModule,
    SimpleNotificationsModule.forRoot(),
    MatMenuModule,
    MatDialogModule,
    Ng2PageScrollModule,
    MatTooltipModule,
    ReactiveFormsModule,
    MatListModule,
    MatExpansionModule,
    MatCheckboxModule,
  ],
  providers: [
    HttpClient,
    ConstService,
    SessionService,
    LoginDialogService,
    UserService,
    CreateProjectDialogService,
    CreateIssueDialogService,
    ChargeService,
    ProjectService,
    TokenService
  ],
  bootstrap: [AppComponent],
  entryComponents: [
    CreateProjectComponent,
    CreateProjectDialogComponent,
    CreateIssueComponent,
    LoginComponent
  ]
})
export class AppModule {}
