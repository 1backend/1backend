"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6022],{13627:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>y,contentTitle:()=>m,default:()=>P,frontMatter:()=>k,metadata:()=>s,toc:()=>f});const s=JSON.parse('{"id":"1backend/get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","source":"@site/docs/1backend/get-public-key.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-public-key","permalink":"/docs/1backend/get-public-key","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","sidebar_label":"Get Public Key","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VMtu2zAQ/BViz7LkJC3Q6lQHKII80ARwgh4cH2hqLTGhSGZJ2VUFAf2IfmG/pFjJzqu9NIeeRFDD5ezsDDsoMCjSPmpnIYcTjCJWKHyzMlqJe2xFdMJLCiikLcQGSa/bAXL29TqFBJxHknz6tIAcSoxXw9FzbCEBwuCdDRgg7+BwOuXPywsvzyEB5WxEG/mv9N5oNRTM7gJDOgiqwlryyhNfF/VY0D/elHcQW4+QQ4ikbQl9n+x33OoOVYS+5713f+NwajfS6EKczS+/CEei1iFoWwqPNCydFbp4O00kcvQvFA/+pHhjZRMrR/o7/i8mvCnLAPkCbgKSmG8ULBOoMVZuN2pIwMtYQQ5ZE5AmYaOycSiT+2H+AWmDxDU6aMgwMDNOSVO5EPP3Hw6PDqBfMk41pGM7Z+Ij12OUhDRruPxrNa5bj+J2B7kFsXbGuC0WYtUKKYKXanRrdPdohVQPjSYsxJpcPVh334+4cKW2Am3hnbaR7ay5foWyQIIErKxZldlO+0FleBRLes3mY6W0XTvmyYORahgM1lJzx0EaDJ/YUI2RkZxNlauf1b46FfPGe0cs5yhSFaPPs+xgJdU92oIPZNAnr1SYCeXq2lmxg4m1I9G6hsTsVDyzRfj142etFTmehlYYJisZsBgUWjXaRI54UNIgC2C0QhuQO9gzPLm6EJujdPqCX8izbLvdpqVtUkdltjsXMll6MzlKp2kVa8OsI1IdLtfz8fan9sJWliVSql02QDJWVkfDkIPjsSdIgA00NjxNj9LphFR6+JHrehdiLe0zpvx6jc+PGN+fF4J1T7F5wzu3m3nEbzHzRmrLFAY5ul0IFrAPAefiKQbLBNjuDOg6lv6GTN/z9kOD1EK+WCawkaTlintfLPtk70DODZfIYaYUenbIRppmNN+r3PfPw3ny+Rr6/jdjDQO1","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Permissions","permalink":"/docs/1backend/save-permissions"},"next":{"title":"Register","permalink":"/docs/1backend/register"}}');var n=i(74848),a=i(28453),l=i(53746),c=i.n(l),r=i(56518),o=i.n(r),d=i(99972),p=i.n(d),u=i(25342),b=i.n(u),h=(i(44215),i(82223),i(24861));const k={id:"get-public-key",title:"Get Public Key",description:"Get the public key to parse and verify the JWT.",sidebar_label:"Get Public Key",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VMtu2zAQ/BViz7LkJC3Q6lQHKII80ARwgh4cH2hqLTGhSGZJ2VUFAf2IfmG/pFjJzqu9NIeeRFDD5ezsDDsoMCjSPmpnIYcTjCJWKHyzMlqJe2xFdMJLCiikLcQGSa/bAXL29TqFBJxHknz6tIAcSoxXw9FzbCEBwuCdDRgg7+BwOuXPywsvzyEB5WxEG/mv9N5oNRTM7gJDOgiqwlryyhNfF/VY0D/elHcQW4+QQ4ikbQl9n+x33OoOVYS+5713f+NwajfS6EKczS+/CEei1iFoWwqPNCydFbp4O00kcvQvFA/+pHhjZRMrR/o7/i8mvCnLAPkCbgKSmG8ULBOoMVZuN2pIwMtYQQ5ZE5AmYaOycSiT+2H+AWmDxDU6aMgwMDNOSVO5EPP3Hw6PDqBfMk41pGM7Z+Ij12OUhDRruPxrNa5bj+J2B7kFsXbGuC0WYtUKKYKXanRrdPdohVQPjSYsxJpcPVh334+4cKW2Am3hnbaR7ay5foWyQIIErKxZldlO+0FleBRLes3mY6W0XTvmyYORahgM1lJzx0EaDJ/YUI2RkZxNlauf1b46FfPGe0cs5yhSFaPPs+xgJdU92oIPZNAnr1SYCeXq2lmxg4m1I9G6hsTsVDyzRfj142etFTmehlYYJisZsBgUWjXaRI54UNIgC2C0QhuQO9gzPLm6EJujdPqCX8izbLvdpqVtUkdltjsXMll6MzlKp2kVa8OsI1IdLtfz8fan9sJWliVSql02QDJWVkfDkIPjsSdIgA00NjxNj9LphFR6+JHrehdiLe0zpvx6jc+PGN+fF4J1T7F5wzu3m3nEbzHzRmrLFAY5ul0IFrAPAefiKQbLBNjuDOg6lv6GTN/z9kOD1EK+WCawkaTlintfLPtk70DODZfIYaYUenbIRppmNN+r3PfPw3ny+Rr6/jdjDQO1",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,y={},f=[];function j(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Get Public Key"}),"\n",(0,n.jsx)(c(),{method:"get",path:"/user-svc/public-key",context:"endpoint"}),"\n",(0,n.jsx)(t.p,{children:"Get the public key to parse and verify the JWT."}),"\n",(0,n.jsx)(o(),{parameters:void 0}),"\n",(0,n.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{publicKey:{type:"string"}},type:"object"}}}},400:{description:"Invalid JSON or missing permission id",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function P(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,n.jsx)(t,{...e,children:(0,n.jsx)(j,{...e})}):j(e)}}}]);