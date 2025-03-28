"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7884],{36719:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>k,contentTitle:()=>m,default:()=>R,frontMatter:()=>h,metadata:()=>i,toc:()=>f});const i=JSON.parse('{"id":"1backend/read-user-by-token","title":"Read User by Token","description":"Retrieve user information based on an authentication token.","source":"@site/docs/1backend/read-user-by-token.api.mdx","sourceDirName":"1backend","slug":"/1backend/read-user-by-token","permalink":"/docs/1backend/read-user-by-token","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"read-user-by-token","title":"Read User by Token","description":"Retrieve user information based on an authentication token.","sidebar_label":"Read User by Token","hide_title":true,"hide_table_of_contents":true,"api":"eJzNVt9v4zYM/lcEPW2AY6fXHTD4ae1+dtehRXN9ygU4RmYctbKkUnJynuH/faDtpGmTDtiGDXtpDZEiP378SKWVBQZF2kftrMzlHUbSuEFRBySh7cpRBWwTSwhYCGcFWAF1XKONWg2m6B7RpjKRziP1R1eFzCUhFPcB6bL5yA4ykYTBOxswyLyV76ZT/vcy/80HmUjlbEQb2QremzFN9hDYpZVBrbEC/vLEGaMeAoKKeoM3VILVv+9htDI2HmUuQyRtS9kl0h249Dd1xCocB1SEELG4iCejFGjwbas+ndpChcdV/1QbI9gk3ErENYpDhDKR+AUqbzjShapQfO/IO9oZj3IEU5fHOe7vricr0mgL04ja6qcaxVfaBl1gn3KmbVkbIGeFNxC58V8LXXCbVxpJrBz1fp8PsX1OX6ADVeFE/Tm62hdvktqxRJ5qTVjIfM4cjoyNRS328dzyAVXkeOMBEEHTxw9I/9dWvmTrV7AofnB4iiUPIWwdFb9AWJ/M/i83mVl81dwHsDgpTsP9q019o5vd8QmffXNqU/Q7RfymQ+Asf3tpIJGj06hPQXl/CsqVjUgWjJghbZDEj33M/wYSSwFVTTo2Mp+38hKBkC7quJb5fNExx1AGJp5XsZhtFPNeYVw73tHehShZbuwvM277JGxU/5Etm0kcN3foKwt9ipoM+2bGKTBrF2L+/tt352eSk+2wzLi2oZxDREdNbDyKT6PLJylWzhi3xUIsGwEieFAowBbDCyNADTISK3JVr9NdTeLaldoKtIV32kYWrub4a4QCabdFcskoHD0v1t3y8PoDNj3D/OT1C8PZCKrvHVagueIABsN3YRihSM6mylUHsW+vxKz23hEzOpC0jtHnWXa2BPWItuAL2bBwDlm4uJpY4LdLVFqRY661wrAfUS7HaIU29Atnl+/n22uxOU+nL7KFPMu2221a2jp1VGbjvZBB6c3kPJ2m61iZfnEiVeFmNRuyPYMNWyhLpFS7rHfJmCcd+x1wdjkUIhPJchjgT9PzdDohlZ5P+9XlQqzAHiC9QyiGVi0bsfst8IKC9nlW/uFPkLGjEb/EzBvQtl9PTE87qnwudyqX42uRyL3SF4lkRbNX23K+ezJdx8dPNRKP2CKRGyANSyZkvuiSnch4NB6x6R9phZ5FsAFTD/p6Nf3d4Qze3sw+ykTCOCHPmuRoye6Dw+9MtjkI/lrTAwb+2yVvXGnbQfFdt/cfTG/e2A/S4M20Lrqu+wPir4fi","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Assign Role","permalink":"/docs/1backend/assign-role"},"next":{"title":"List Users","permalink":"/docs/1backend/get-users"}}');var r=n(74848),s=n(28453),o=n(53746),a=n.n(o),d=n(56518),p=n.n(d),c=n(99972),l=n.n(c),b=n(25342),u=n.n(b),y=(n(44215),n(82223),n(24861));const h={id:"read-user-by-token",title:"Read User by Token",description:"Retrieve user information based on an authentication token.",sidebar_label:"Read User by Token",hide_title:!0,hide_table_of_contents:!0,api:"eJzNVt9v4zYM/lcEPW2AY6fXHTD4ae1+dtehRXN9ygU4RmYctbKkUnJynuH/faDtpGmTDtiGDXtpDZEiP378SKWVBQZF2kftrMzlHUbSuEFRBySh7cpRBWwTSwhYCGcFWAF1XKONWg2m6B7RpjKRziP1R1eFzCUhFPcB6bL5yA4ykYTBOxswyLyV76ZT/vcy/80HmUjlbEQb2QremzFN9hDYpZVBrbEC/vLEGaMeAoKKeoM3VILVv+9htDI2HmUuQyRtS9kl0h249Dd1xCocB1SEELG4iCejFGjwbas+ndpChcdV/1QbI9gk3ErENYpDhDKR+AUqbzjShapQfO/IO9oZj3IEU5fHOe7vricr0mgL04ja6qcaxVfaBl1gn3KmbVkbIGeFNxC58V8LXXCbVxpJrBz1fp8PsX1OX6ADVeFE/Tm62hdvktqxRJ5qTVjIfM4cjoyNRS328dzyAVXkeOMBEEHTxw9I/9dWvmTrV7AofnB4iiUPIWwdFb9AWJ/M/i83mVl81dwHsDgpTsP9q019o5vd8QmffXNqU/Q7RfymQ+Asf3tpIJGj06hPQXl/CsqVjUgWjJghbZDEj33M/wYSSwFVTTo2Mp+38hKBkC7quJb5fNExx1AGJp5XsZhtFPNeYVw73tHehShZbuwvM277JGxU/5Etm0kcN3foKwt9ipoM+2bGKTBrF2L+/tt352eSk+2wzLi2oZxDREdNbDyKT6PLJylWzhi3xUIsGwEieFAowBbDCyNADTISK3JVr9NdTeLaldoKtIV32kYWrub4a4QCabdFcskoHD0v1t3y8PoDNj3D/OT1C8PZCKrvHVagueIABsN3YRihSM6mylUHsW+vxKz23hEzOpC0jtHnWXa2BPWItuAL2bBwDlm4uJpY4LdLVFqRY661wrAfUS7HaIU29Atnl+/n22uxOU+nL7KFPMu2221a2jp1VGbjvZBB6c3kPJ2m61iZfnEiVeFmNRuyPYMNWyhLpFS7rHfJmCcd+x1wdjkUIhPJchjgT9PzdDohlZ5P+9XlQqzAHiC9QyiGVi0bsfst8IKC9nlW/uFPkLGjEb/EzBvQtl9PTE87qnwudyqX42uRyL3SF4lkRbNX23K+ezJdx8dPNRKP2CKRGyANSyZkvuiSnch4NB6x6R9phZ5FsAFTD/p6Nf3d4Qze3sw+ykTCOCHPmuRoye6Dw+9MtjkI/lrTAwb+2yVvXGnbQfFdt/cfTG/e2A/S4M20Lrqu+wPir4fi",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,k={},f=[];function j(e){const t={p:"p",...(0,s.R)(),...e.components};return(0,r.jsxs)(r.Fragment,{children:[(0,r.jsx)(y.default,{as:"h1",className:"openapi__heading",children:"Read User by Token"}),"\n",(0,r.jsx)(a(),{method:"post",path:"/user-svc/user/by-token",context:"endpoint"}),"\n",(0,r.jsx)(t.p,{children:"Retrieve user information based on an authentication token."}),"\n",(0,r.jsx)(p(),{parameters:void 0}),"\n",(0,r.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,r.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{activeOrganizationId:{type:"string"},organizations:{items:{properties:{createdAt:{type:"string"},deletedAt:{type:"string"},id:{type:"string"},name:{description:"Full name of the organization",example:"Acme Corporation",type:"string"},slug:{description:"URL-friendly unique (inside the Singularon platform) identifier for the `organization`.",example:"acme-corporation",type:"string"},updatedAt:{type:"string"}},required:["id","name","slug"],type:"object"},type:"array"},user:{properties:{createdAt:{type:"string"},deletedAt:{type:"string"},id:{type:"string"},name:{description:"Full name of the organization.",example:"Jane Doe",type:"string"},passwordHash:{type:"string"},slug:{description:"URL-friendly unique (inside the Singularon platform) identifier for the `user`.",example:"jane-doe",type:"string"},updatedAt:{type:"string"}},required:["id","slug"],type:"object"}},type:"object"}}}},400:{description:"Token Missing",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function R(e={}){const{wrapper:t}={...(0,s.R)(),...e.components};return t?(0,r.jsx)(t,{...e,children:(0,r.jsx)(j,{...e})}):j(e)}}}]);