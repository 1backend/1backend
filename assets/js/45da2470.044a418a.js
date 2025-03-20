"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[8696],{5466:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>k,contentTitle:()=>y,default:()=>x,frontMatter:()=>b,metadata:()=>i,toc:()=>m});const i=JSON.parse('{"id":"1backend/download-file","title":"Download a File","description":"Start or resume the download for a specified URL.","source":"@site/docs/1backend/download-file.api.mdx","sourceDirName":"1backend","slug":"/1backend/download-file","permalink":"/docs/1backend/download-file","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"download-file","title":"Download a File","description":"Start or resume the download for a specified URL.","sidebar_label":"Download a File","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VVFv4zYM/isCnx07bXfA5qe1223oVuyC5vKUBjhGph1dZUmV5OSywP99oB3nnF5u2F76ktgWSX38yI88QEFBeuWisgZymEf0UVgvPIWmJhE3JAq7M9piIUrrBYrgSKpSUSEWjw/pk3kyj/TSKE+hs/5UKk2TsJX54JdLTxjpk3DkaxWCsiaFBKwjj3ztfQE5DMa/KU2QgKeXhkK8s8Ue8gNIayKZyI/onFayc8w+BwZ9gCA3VCM/Oc9ho6LAb6XVBfkZxg2/xb0jyCFEr0wFbQKN1xe+t/3tylMB+bIzWiWDkV1/JhmhZatz5n4dWHrsocM4TPQNdXGDsyb04K6nU/77ThRlVFQYqRChkZJCKBut95D8dyqwKBQfoZ6NSGEkF7JpE/jhEp57s0WtCvHH/MNf/+fy8zqQ99ZfpvoykqtvkSwMNnFjvfqbijdD8u4yJ5G8QS3m5Lfkxfsu5ttAahMIJBuv4h7y5QHuCD3524Y7fLlquVGxCty3LCQx30pu3prixrLKXMN96TpBQDYoNRvEBxydcwpd8E4fkGXaStQbG2L+7sfrmyvgawYUc86qT2SM5TVnH/eOxNPR5AlEabW2OyrEet+NFJQk0BQi2mcyAmUvHFF6W3dTZRHIczbiwVbKCDKFs8pEniOK428IC+IiGKyZsttjr3QlgBOT6NSftO+45bI8fp0y779g7TS9nhpDNY7DYjQ8lCntMJpQdlWnGlVnhJrCz0GZqtEYvTWptPUI2+xezBvnrI+nuJsYXZ5lV2uUz2QKdsjgmxlzez8xGNWWRK2kt1wrJSkIpzGW1tdMh1aSTOgSGe77ffYgtjfp9Oy2kGfZbrdLK9Ok1lfZ0S9kWDk9uUmn6SbWmjFE8nX4UM77276CDTusKvKpsllnkjHPKjKJcHXXJwIJcDv18KfpTTqdeJle/8RxnQ2xRjNCepp+KI574Cz/0R54y0V17J1IX2LmNCoz2h29lJYwxGLIg5hWCbBo+PhwWGOghddty59fGvKs31UCW/QK18zZctUmQx+z+p5pDzn80mc8YQExmaibvpFfDZg2GTxupSQX/9V2PBJmi4+QwPq4aGtbsIvHHe8v3EEO3apm507k3bcDaDRVgxXb9iFZUticLVp0ihElwwMnNRyZ/Qjga7H2efAvZ3XR5XDopdy2J/v+6LsepwnRW3MVV23b/gOzFx2G","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Send an Email","permalink":"/docs/1backend/send-email"},"next":{"title":"Get a Download","permalink":"/docs/1backend/get-download"}}');var a=n(74848),o=n(28453),s=n(53746),d=n.n(s),r=n(56518),p=n.n(r),l=n(99972),c=n.n(l),f=n(25342),h=n.n(f),u=(n(44215),n(82223),n(24861));const b={id:"download-file",title:"Download a File",description:"Start or resume the download for a specified URL.",sidebar_label:"Download a File",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VVFv4zYM/isCnx07bXfA5qe1223oVuyC5vKUBjhGph1dZUmV5OSywP99oB3nnF5u2F76ktgWSX38yI88QEFBeuWisgZymEf0UVgvPIWmJhE3JAq7M9piIUrrBYrgSKpSUSEWjw/pk3kyj/TSKE+hs/5UKk2TsJX54JdLTxjpk3DkaxWCsiaFBKwjj3ztfQE5DMa/KU2QgKeXhkK8s8Ue8gNIayKZyI/onFayc8w+BwZ9gCA3VCM/Oc9ho6LAb6XVBfkZxg2/xb0jyCFEr0wFbQKN1xe+t/3tylMB+bIzWiWDkV1/JhmhZatz5n4dWHrsocM4TPQNdXGDsyb04K6nU/77ThRlVFQYqRChkZJCKBut95D8dyqwKBQfoZ6NSGEkF7JpE/jhEp57s0WtCvHH/MNf/+fy8zqQ99ZfpvoykqtvkSwMNnFjvfqbijdD8u4yJ5G8QS3m5Lfkxfsu5ttAahMIJBuv4h7y5QHuCD3524Y7fLlquVGxCty3LCQx30pu3prixrLKXMN96TpBQDYoNRvEBxydcwpd8E4fkGXaStQbG2L+7sfrmyvgawYUc86qT2SM5TVnH/eOxNPR5AlEabW2OyrEet+NFJQk0BQi2mcyAmUvHFF6W3dTZRHIczbiwVbKCDKFs8pEniOK428IC+IiGKyZsttjr3QlgBOT6NSftO+45bI8fp0y779g7TS9nhpDNY7DYjQ8lCntMJpQdlWnGlVnhJrCz0GZqtEYvTWptPUI2+xezBvnrI+nuJsYXZ5lV2uUz2QKdsjgmxlzez8xGNWWRK2kt1wrJSkIpzGW1tdMh1aSTOgSGe77ffYgtjfp9Oy2kGfZbrdLK9Ok1lfZ0S9kWDk9uUmn6SbWmjFE8nX4UM77276CDTusKvKpsllnkjHPKjKJcHXXJwIJcDv18KfpTTqdeJle/8RxnQ2xRjNCepp+KI574Cz/0R54y0V17J1IX2LmNCoz2h29lJYwxGLIg5hWCbBo+PhwWGOghddty59fGvKs31UCW/QK18zZctUmQx+z+p5pDzn80mc8YQExmaibvpFfDZg2GTxupSQX/9V2PBJmi4+QwPq4aGtbsIvHHe8v3EEO3apm507k3bcDaDRVgxXb9iFZUticLVp0ihElwwMnNRyZ/Qjga7H2efAvZ3XR5XDopdy2J/v+6LsepwnRW3MVV23b/gOzFx2G",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},y=void 0,k={},m=[];function w(e){const t={code:"code",p:"p",...(0,o.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"Download a File"}),"\n",(0,a.jsx)(d(),{method:"put",path:"/file-svc/download",context:"endpoint"}),"\n",(0,a.jsx)(t.p,{children:"Start or resume the download for a specified URL."}),"\n",(0,a.jsxs)(t.p,{children:["Requires the ",(0,a.jsx)(t.code,{children:"file-svc:download:create"})," permission."]}),"\n",(0,a.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(p(),{parameters:void 0}),"\n",(0,a.jsx)(c(),{title:"Body",body:{content:{"application/json":{schema:{properties:{folderPath:{type:"string"},url:{type:"string"}},required:["url"],type:"object"}}},description:"Download Request",required:!0}}),"\n",(0,a.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"Download initiated successfully",content:{"application/json":{schema:{additionalProperties:!0,type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function x(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(w,{...e})}):w(e)}}}]);