"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9517],{62140:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>k,contentTitle:()=>j,default:()=>m,frontMatter:()=>u,metadata:()=>n,toc:()=>w});const n=JSON.parse('{"id":"1backend/change-password","title":"Change User Password","description":"Allows an authenticated user to change their own password.","source":"@site/docs/1backend/change-password.api.mdx","sourceDirName":"1backend","slug":"/1backend/change-password","permalink":"/docs/1backend/change-password","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"change-password","title":"Change User Password","description":"Allows an authenticated user to change their own password.","sidebar_label":"Change User Password","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU2P2zYQ/SvEnGXJrhsU0KneIgi2DbqLOHtyfBhTY4lZimRIyo4r6L8XI1m21+sGKQLkYgvkfLx5M2/YQkFBeuWisgZyWGht90GgEdjEikxUEiMVognkRbRCVmhKErEi5YXdG+EwhL31RQoJWEceOdB9ATkMpo/He0jA05eGQryzxQHyFqQ1kUzkT3ROcyJlTfY5MJAWgqyoRv5yngNHRaF3a7wnE09x8xbiwRHkEKJXpoQuAUP7b94H3ZQ3LrpkPLGbzyQjdHz0kqA/BgLG8OLDUNSxPOWpgDz6hjo+CM6aMMD+ZTrlv5fBTlEGrgoRGikphG2j9QGS76foFewugV9vZbw3O9SqEH8uH/7+Pwle9oC8t/47CeyRzF4jeTI8Ydarf6j4aUje3OYkkjeoxZL8jrx428f8OZB4GEk2XsUD5KsW7gg9+UUTK8hX626dQMQyQL6CJ1bgcidhnUBNsbKsMWf70XPI9pCxSidhJ7NhnibuLL7Q1xb6JI3XbJ1pK1FXNsR8NpvPfwNON6JZcnVDQZeYrrn7eHAkPh1NPoHYWl4gVIjNQaAIDiUJNIWI9pmMQDlIRGy9rXmJiLEq8d6WyggyhbPKRN4miuNXhAVxMwzWTN3iODN9K+DEKDr1Fx16jrk9H86b5u1XrJ2mm5tjbM3VwjgfD3viYnEos7Xj7kLZjwbVqJjOgJrC70GZstEYvTWptPUF8Md7sWycs54bNnSgitHlWTbboHwmU7BDBq82zuJ+YjCqHYlaSW+5kUpSEE5j3FpfM1daSTKhr3LM9+7xvdjN0+mLbCHPsv1+n5amSa0vs6NfyLB0ejJPp2kVa80YIvk6PGyXQ7Yz2LDHsiSfKpv1Jhk3QUVmGGZ3QyGQAM/aAH+aztPpxMt0/obj8sTWaC6QHjdqPwkXr8ULEi5eix97oo7zEulrzJxGZRhTz097VNEKRhXxCrjS0ToB1gtbte0GAz153XV8/KUhzxJeJ7BDr3DDjKzWXTKOMAvvmQ5c8FDKhLXDVKFuhhm+2jFdMnospCQXv2l7uRUeH5YfIYHN8aGtbcE+Hvf8SuEecugfa/buBd6ftaDRlA2WbDvEZDnhUfZnoTGkZPzgqsYrc7hAeC3UoRD+5bJuurTtIOOuO9kPV//pcdoOgzV3c9113b8O6A6t","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Reset Password","permalink":"/docs/1backend/reset-password"},"next":{"title":"List Enrolls","permalink":"/docs/1backend/list-enrolls"}}');var a=t(74848),r=t(28453),o=t(53746),i=t.n(o),d=t(56518),p=t.n(d),c=t(99972),h=t.n(c),l=t(25342),b=t.n(l),g=(t(44215),t(82223),t(24861));const u={id:"change-password",title:"Change User Password",description:"Allows an authenticated user to change their own password.",sidebar_label:"Change User Password",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU2P2zYQ/SvEnGXJrhsU0KneIgi2DbqLOHtyfBhTY4lZimRIyo4r6L8XI1m21+sGKQLkYgvkfLx5M2/YQkFBeuWisgZyWGht90GgEdjEikxUEiMVognkRbRCVmhKErEi5YXdG+EwhL31RQoJWEceOdB9ATkMpo/He0jA05eGQryzxQHyFqQ1kUzkT3ROcyJlTfY5MJAWgqyoRv5yngNHRaF3a7wnE09x8xbiwRHkEKJXpoQuAUP7b94H3ZQ3LrpkPLGbzyQjdHz0kqA/BgLG8OLDUNSxPOWpgDz6hjo+CM6aMMD+ZTrlv5fBTlEGrgoRGikphG2j9QGS76foFewugV9vZbw3O9SqEH8uH/7+Pwle9oC8t/47CeyRzF4jeTI8Ydarf6j4aUje3OYkkjeoxZL8jrx428f8OZB4GEk2XsUD5KsW7gg9+UUTK8hX626dQMQyQL6CJ1bgcidhnUBNsbKsMWf70XPI9pCxSidhJ7NhnibuLL7Q1xb6JI3XbJ1pK1FXNsR8NpvPfwNON6JZcnVDQZeYrrn7eHAkPh1NPoHYWl4gVIjNQaAIDiUJNIWI9pmMQDlIRGy9rXmJiLEq8d6WyggyhbPKRN4miuNXhAVxMwzWTN3iODN9K+DEKDr1Fx16jrk9H86b5u1XrJ2mm5tjbM3VwjgfD3viYnEos7Xj7kLZjwbVqJjOgJrC70GZstEYvTWptPUF8Md7sWycs54bNnSgitHlWTbboHwmU7BDBq82zuJ+YjCqHYlaSW+5kUpSEE5j3FpfM1daSTKhr3LM9+7xvdjN0+mLbCHPsv1+n5amSa0vs6NfyLB0ejJPp2kVa80YIvk6PGyXQ7Yz2LDHsiSfKpv1Jhk3QUVmGGZ3QyGQAM/aAH+aztPpxMt0/obj8sTWaC6QHjdqPwkXr8ULEi5eix97oo7zEulrzJxGZRhTz097VNEKRhXxCrjS0ToB1gtbte0GAz153XV8/KUhzxJeJ7BDr3DDjKzWXTKOMAvvmQ5c8FDKhLXDVKFuhhm+2jFdMnospCQXv2l7uRUeH5YfIYHN8aGtbcE+Hvf8SuEecugfa/buBd6ftaDRlA2WbDvEZDnhUfZnoTGkZPzgqsYrc7hAeC3UoRD+5bJuurTtIOOuO9kPV//pcdoOgzV3c9113b8O6A6t",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,k={},w=[];function f(e){const s={p:"p",...(0,r.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Change User Password"}),"\n",(0,a.jsx)(i(),{method:"post",path:"/user-svc/change-password",context:"endpoint"}),"\n",(0,a.jsx)(s.p,{children:"Allows an authenticated user to change their own password."}),"\n",(0,a.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(p(),{parameters:void 0}),"\n",(0,a.jsx)(h(),{title:"Body",body:{content:{"application/json":{schema:{properties:{currentPassword:{type:"string"},newPassword:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"Change Password Request",required:!0}}),"\n",(0,a.jsx)(b(),{id:void 0,label:void 0,responses:{200:{description:"Password changed successfully",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function m(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,a.jsx)(s,{...e,children:(0,a.jsx)(f,{...e})}):f(e)}}}]);