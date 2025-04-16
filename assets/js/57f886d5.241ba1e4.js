"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[6022],{13627:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>g,contentTitle:()=>j,default:()=>K,frontMatter:()=>u,metadata:()=>i,toc:()=>y});const i=JSON.parse('{"id":"1backend/get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","source":"@site/docs/1backend/get-public-key.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-public-key","permalink":"/docs/1backend/get-public-key","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","sidebar_label":"Get Public Key","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE1v2zAM/SsCz46dLAM2+LQUGIq0xVogLXZoc1BkxlYjSxolO8sM//eBdtKv9dTDThbkR/Lx8VEdFBgUaR+1s5DDOUYRKxS+2RitxA4PIjrhJQUU0haiRdLbwwC5+HmbQgLOI0mOXhaQQ4nxZgi9xAMkQBi8swED5B18mk7587rg9SUkoJyNaCP/ld4brYaE2WNgSAdBVVhLPnniclGPCf1TpbyDePAIOYRI2pbQ91z8V6MJC8jvX0DXyQnqNo+oIvQ9gz+/R25pW2l0IS5W1z+EI1HrELQthUcajs4KXXycPxI5ep/7+xRn/1K8s7KJlSP9B/8XE76UZWBZ7wKSWLWKVa0xVu7oAUjAy1hBDlkTkCahVdk4gsluMEZAapE4RwcNGQZmxilpKhdiPpvN51+gXzNONaTjYcXER65nKAlp0XD6t2rcHjyKhyPkAcTWGeP2WIjNQUgRvFSjjaPboRVSjQYRW3L14OlTP+LKldoKtIV32kb2ueb8FcoCCRKwsmZVFkftB5XhSSzpNVuNldJ265gnD0aqYTBYS80dB2kwfGNDNUZGcjZVrn6R+2YpVo33jljOUaQqRp9n2Wwj1Q5twQEZ9MkbFRbLiZVRtyhqrcix1lphEN7IuHVUcztGK7QBmc+p3vnNlWjn6fRVtZBn2X6/T0vbpI7K7BgXMll6M5mn07SKtWEOEakO19vVWO2ZbNjLskRKtcsGSMY66WgYMjsbG4EE2A4j/Wk6T6cTUun8K+f1LsRa2hdM+ZEaXxkxPjOv2u+el+ADz9lxghF/x8wbqS1TGOTojpa+h5Ol2eXPpl4nwOZlQNdtZMA7Mn3P178apAPk9+sEWklabrj3+3WfnPzEW8ApclgohZ7n3UrTjFZ6s8X9y1U7/34Lff8X5aD5Ww==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Permits","permalink":"/docs/1backend/save-permits"},"next":{"title":"Register","permalink":"/docs/1backend/register"}}');var a=n(74848),s=n(28453),l=n(53746),c=n.n(l),p=n(56518),o=n.n(p),r=n(99972),d=n.n(r),b=n(25342),k=n.n(b),h=(n(44215),n(82223),n(24861));const u={id:"get-public-key",title:"Get Public Key",description:"Get the public key to parse and verify the JWT.",sidebar_label:"Get Public Key",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE1v2zAM/SsCz46dLAM2+LQUGIq0xVogLXZoc1BkxlYjSxolO8sM//eBdtKv9dTDThbkR/Lx8VEdFBgUaR+1s5DDOUYRKxS+2RitxA4PIjrhJQUU0haiRdLbwwC5+HmbQgLOI0mOXhaQQ4nxZgi9xAMkQBi8swED5B18mk7587rg9SUkoJyNaCP/ld4brYaE2WNgSAdBVVhLPnniclGPCf1TpbyDePAIOYRI2pbQ91z8V6MJC8jvX0DXyQnqNo+oIvQ9gz+/R25pW2l0IS5W1z+EI1HrELQthUcajs4KXXycPxI5ep/7+xRn/1K8s7KJlSP9B/8XE76UZWBZ7wKSWLWKVa0xVu7oAUjAy1hBDlkTkCahVdk4gsluMEZAapE4RwcNGQZmxilpKhdiPpvN51+gXzNONaTjYcXER65nKAlp0XD6t2rcHjyKhyPkAcTWGeP2WIjNQUgRvFSjjaPboRVSjQYRW3L14OlTP+LKldoKtIV32kb2ueb8FcoCCRKwsmZVFkftB5XhSSzpNVuNldJ265gnD0aqYTBYS80dB2kwfGNDNUZGcjZVrn6R+2YpVo33jljOUaQqRp9n2Wwj1Q5twQEZ9MkbFRbLiZVRtyhqrcix1lphEN7IuHVUcztGK7QBmc+p3vnNlWjn6fRVtZBn2X6/T0vbpI7K7BgXMll6M5mn07SKtWEOEakO19vVWO2ZbNjLskRKtcsGSMY66WgYMjsbG4EE2A4j/Wk6T6cTUun8K+f1LsRa2hdM+ZEaXxkxPjOv2u+el+ADz9lxghF/x8wbqS1TGOTojpa+h5Ol2eXPpl4nwOZlQNdtZMA7Mn3P178apAPk9+sEWklabrj3+3WfnPzEW8ApclgohZ7n3UrTjFZ6s8X9y1U7/34Lff8X5aD5Ww==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},j=void 0,g={},y=[];function x(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Get Public Key"}),"\n",(0,a.jsx)(c(),{method:"get",path:"/user-svc/public-key",context:"endpoint"}),"\n",(0,a.jsx)(t.p,{children:"Get the public key to parse and verify the JWT."}),"\n",(0,a.jsx)(o(),{parameters:void 0}),"\n",(0,a.jsx)(d(),{title:"Body",body:void 0}),"\n",(0,a.jsx)(k(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{publicKey:{type:"string"}},required:["publicKey"],type:"object"}}}},400:{description:"Invalid JSON or missing permission id",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function K(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(x,{...e})}):x(e)}}}]);