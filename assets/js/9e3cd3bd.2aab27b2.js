"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[702],{63105:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>v,contentTitle:()=>k,default:()=>E,frontMatter:()=>m,metadata:()=>s,toc:()=>b});const s=JSON.parse('{"id":"1backend/events","title":"Events","description":"Events is a dummy endpoint to display documentation about the events that this service emits.","source":"@site/docs/1backend/events.api.mdx","sourceDirName":"1backend","slug":"/1backend/events","permalink":"/docs/1backend/events","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"events","title":"Events","description":"Events is a dummy endpoint to display documentation about the events that this service emits.","sidebar_label":"Events","hide_title":true,"hide_table_of_contents":true,"api":"eJy1VE1P3DAQ/SvWnEOygCq1ORUqhBBIIAEn2MOsM5sYHNvYk92mUaT+iP7C/pJqkl2+7s0liT0zfu/NGw9QUdLRBDbeQQlnG3KclEkKVdW1ba/IVcEbx4q9qkwKFntVed215BglS+HKd6y4IUVzNjco/yapRHFjNClqDaccMvCB4pR1UUEJczxkECkF7xIlKAc4Wizk9RHY9SVkoL1jciy7GII1eipVPCUJGSDphlqUrxDlIDZzQW4iYSUnDsB9ICghcTSuhnHM9it+9USaYZQnA8Y6QfkAP4TK7UbDMoOWuPECuyaGDAJyAyUUukE+SBtdvLIR1hQlf4AuWgkqrNdoG5+4/PL16PgQxqXE6S4a7m8F9wz1lDBSPOmk9GcJ7vpA6nEX8ghq7a31W6rUqleoUkBNCl2l2D+TU6hfOhOpUuvo26k794micFFXvjbutbHSFiP1G8KKImTgsBVFBIWP5tckMrwKhcFcUj9JZ9zaC07pC+qpL9SiEcYJLaXvybi6s8jRu1z79l3tmwt124Xgo0g5i9Qwh7IoDleon8lVklDAmH1S4UR1zqyN0J7jVEUbsj6IIVWwyGsfW7X2UfW+i+rkQr3zSvr7+09rdPQ7Z6aDFSaqJt1WnbGTzZNGSyKLNZpcIuG1x31+c6U2x/niA+pUFsV2u81r1+U+1sUuLxVYB3twnC/yhlsrXJhim67Xt/Ppb6TTFuuaYm58MYUUordhKyGHpzNRyEBsNcuwyI/zxUHU+dE3qRt84hbdO6Rnezd+kG94m6H/Pus7vzD95CJYNE6ATqINu+F5gP3wQLa/DJYZyJjI5jBIc+6jHUdZfuko9lA+LDPYYDS4EnUelmO2d67M2zP10iatKYizNmi72bSfrovx/UCfn93BOP4DhIq+8g==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Introduction","permalink":"/docs/1backend/1-backend"},"next":{"title":"Delete a Message","permalink":"/docs/1backend/delete-message"}}');var i=n(74848),a=n(28453),d=n(53746),o=n.n(d),c=n(56518),l=n.n(c),r=n(99972),p=n.n(r),u=n(25342),U=n.n(u),h=(n(44215),n(82223),n(24861));const m={id:"events",title:"Events",description:"Events is a dummy endpoint to display documentation about the events that this service emits.",sidebar_label:"Events",hide_title:!0,hide_table_of_contents:!0,api:"eJy1VE1P3DAQ/SvWnEOygCq1ORUqhBBIIAEn2MOsM5sYHNvYk92mUaT+iP7C/pJqkl2+7s0liT0zfu/NGw9QUdLRBDbeQQlnG3KclEkKVdW1ba/IVcEbx4q9qkwKFntVed215BglS+HKd6y4IUVzNjco/yapRHFjNClqDaccMvCB4pR1UUEJczxkECkF7xIlKAc4Wizk9RHY9SVkoL1jciy7GII1eipVPCUJGSDphlqUrxDlIDZzQW4iYSUnDsB9ICghcTSuhnHM9it+9USaYZQnA8Y6QfkAP4TK7UbDMoOWuPECuyaGDAJyAyUUukE+SBtdvLIR1hQlf4AuWgkqrNdoG5+4/PL16PgQxqXE6S4a7m8F9wz1lDBSPOmk9GcJ7vpA6nEX8ghq7a31W6rUqleoUkBNCl2l2D+TU6hfOhOpUuvo26k794micFFXvjbutbHSFiP1G8KKImTgsBVFBIWP5tckMrwKhcFcUj9JZ9zaC07pC+qpL9SiEcYJLaXvybi6s8jRu1z79l3tmwt124Xgo0g5i9Qwh7IoDleon8lVklDAmH1S4UR1zqyN0J7jVEUbsj6IIVWwyGsfW7X2UfW+i+rkQr3zSvr7+09rdPQ7Z6aDFSaqJt1WnbGTzZNGSyKLNZpcIuG1x31+c6U2x/niA+pUFsV2u81r1+U+1sUuLxVYB3twnC/yhlsrXJhim67Xt/Ppb6TTFuuaYm58MYUUordhKyGHpzNRyEBsNcuwyI/zxUHU+dE3qRt84hbdO6Rnezd+kG94m6H/Pus7vzD95CJYNE6ATqINu+F5gP3wQLa/DJYZyJjI5jBIc+6jHUdZfuko9lA+LDPYYDS4EnUelmO2d67M2zP10iatKYizNmi72bSfrovx/UCfn93BOP4DhIq+8g==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},k=void 0,v={},b=[];function y(e){const t={p:"p",...(0,a.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(h.default,{as:"h1",className:"openapi__heading",children:"Events"}),"\n",(0,i.jsx)(o(),{method:"get",path:"/chat-svc/events",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Events is a dummy endpoint to display documentation about the events that this service emits."}),"\n",(0,i.jsx)(l(),{parameters:void 0}),"\n",(0,i.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(U(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{threadId:{type:"string"}},type:"object"}}}}}})]})}function E(e={}){const{wrapper:t}={...(0,a.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(y,{...e})}):y(e)}}}]);