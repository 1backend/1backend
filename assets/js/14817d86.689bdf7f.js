"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3318],{21211:(e,n,s)=>{s.r(n),s.d(n,{assets:()=>m,contentTitle:()=>J,default:()=>G,frontMatter:()=>u,metadata:()=>t,toc:()=>k});const t=JSON.parse('{"id":"1backend/save-self","title":"Save User Profile","description":"Save user\'s own profile information.","source":"@site/docs/1backend/save-self.api.mdx","sourceDirName":"1backend","slug":"/1backend/save-self","permalink":"/docs/1backend/save-self","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"save-self","title":"Save User Profile","description":"Save user\'s own profile information.","sidebar_label":"Save User Profile","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VVFv2zYQ/ivEvWwDFMmZO2zQ05y1G7wWSxY32ENqDGfpLLGhSJak7HiC/vtwohU7sVd0L32xRfLI++47fh87KMkXTtogjYYcFrgh0Xpy33hhtlpYZ9ZSkZB6bVyDHJVCAsaSGwbzEnLwuKEFqTUkYNFhQ4Gch/y+e3H4nScn5q8hAclDi6GGBDQ2BDlw0nkJCTj61EpHJeTBtZSAL2pqEPIOws5ypA9O6gr6fhmDyYcrU+44ojA6kA78idYqWQwgs4+e83dHR1nHJQRJnkcRwkmCBLxqq7MLoW6blUapfpWKmIQO6BEbqziMGft7/efrx9b/9ZOpIDlB/jRjVh+pCNDz1JlO3Ozpv41lnrDT84S3RvtYyPeTCf89P+n6LSRfzswJsD6BV+eOnesNKlmK3xfXf/yfBM+pJ+eMO9fcMxQNSC5PkdxpbENtnPyHyq+G5IfznARyGpVYkNuQE2+GM78OJL6uVLROht2gvStCR27Whhry+yWLJWDFsow6XGwKWCbQUKgNi9i2YdAvh0PGarzwmyLzUdd+qCeKunWKQzJlClS18SG/vJxOfwROMSJYcEWxiGMcL/l6v7MkPuxDPoBYG6XMlkqx2gkU3mJBAnUpgnkgLbCId1+snWlEqEmMlYh3ppJakC6tkTqko8XUhCW5g8nM9vdkoP+gS7TyLe0GXrkltwdTeTOKejSJsQejNxzGJ5Zw4gM9o1qb0aiwGC4ENSjV4KKK/M9e6qpVGJzRaWGaI+g3c7ForTWO+xR7UIdg8yy7XGHxQLrkDRmcOMlsfqExyA2JRhbOcCtlQV5YhYFdndlSsiDtj+uc/XbzTmym6eRZNp9n2Xa7TSvdpsZV2X6fz7Cy6mKaTtI6NGpwSHKNv14vYrYDWL/FqiKXSpMNIRlTJ8NgnJdXsRBIgG9bhD9Jp+nkwhXp9BWfa40PDeojpINTDjdhb5fwgoGjd+FLH7j9xQj0GDKrUGpOPdDQ7TVyD6NGBnmoNauJ1cBLXbdCT3dO9T1Pf2rJsSiXCWzQSVxxtfw8Ss/fJeRrVJ4+A/vb273vfycOr+hZkOOV1jtmEVXLI0jggXaHV7Zf9skoDgYSF3+J6S5YlUebTxyrT8Yds6IgGz4be+wxN3fvIYHV/rFuTMlbHG75XcNtxGmG6gfnGOY6UKirFiuOjUeyTnHvJwcFM6Jk/OCizjLx0gFiHfzLVZ3d0nXRH/r+KT4u/eeOJ9uJ0dyYZd/3/wLD2zml","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register","permalink":"/docs/1backend/register"},"next":{"title":"Has Permission","permalink":"/docs/1backend/has-permission"}}');var i=s(74848),r=s(28453),a=s(53746),o=s.n(a),p=s(56518),d=s.n(p),l=s(99972),c=s.n(l),f=s(25342),h=s.n(f),b=(s(44215),s(82223),s(24861));const u={id:"save-self",title:"Save User Profile",description:"Save user's own profile information.",sidebar_label:"Save User Profile",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VVFv2zYQ/ivEvWwDFMmZO2zQ05y1G7wWSxY32ENqDGfpLLGhSJak7HiC/vtwohU7sVd0L32xRfLI++47fh87KMkXTtogjYYcFrgh0Xpy33hhtlpYZ9ZSkZB6bVyDHJVCAsaSGwbzEnLwuKEFqTUkYNFhQ4Gch/y+e3H4nScn5q8hAclDi6GGBDQ2BDlw0nkJCTj61EpHJeTBtZSAL2pqEPIOws5ypA9O6gr6fhmDyYcrU+44ojA6kA78idYqWQwgs4+e83dHR1nHJQRJnkcRwkmCBLxqq7MLoW6blUapfpWKmIQO6BEbqziMGft7/efrx9b/9ZOpIDlB/jRjVh+pCNDz1JlO3Ozpv41lnrDT84S3RvtYyPeTCf89P+n6LSRfzswJsD6BV+eOnesNKlmK3xfXf/yfBM+pJ+eMO9fcMxQNSC5PkdxpbENtnPyHyq+G5IfznARyGpVYkNuQE2+GM78OJL6uVLROht2gvStCR27Whhry+yWLJWDFsow6XGwKWCbQUKgNi9i2YdAvh0PGarzwmyLzUdd+qCeKunWKQzJlClS18SG/vJxOfwROMSJYcEWxiGMcL/l6v7MkPuxDPoBYG6XMlkqx2gkU3mJBAnUpgnkgLbCId1+snWlEqEmMlYh3ppJakC6tkTqko8XUhCW5g8nM9vdkoP+gS7TyLe0GXrkltwdTeTOKejSJsQejNxzGJ5Zw4gM9o1qb0aiwGC4ENSjV4KKK/M9e6qpVGJzRaWGaI+g3c7ForTWO+xR7UIdg8yy7XGHxQLrkDRmcOMlsfqExyA2JRhbOcCtlQV5YhYFdndlSsiDtj+uc/XbzTmym6eRZNp9n2Xa7TSvdpsZV2X6fz7Cy6mKaTtI6NGpwSHKNv14vYrYDWL/FqiKXSpMNIRlTJ8NgnJdXsRBIgG9bhD9Jp+nkwhXp9BWfa40PDeojpINTDjdhb5fwgoGjd+FLH7j9xQj0GDKrUGpOPdDQ7TVyD6NGBnmoNauJ1cBLXbdCT3dO9T1Pf2rJsSiXCWzQSVxxtfw8Ss/fJeRrVJ4+A/vb273vfycOr+hZkOOV1jtmEVXLI0jggXaHV7Zf9skoDgYSF3+J6S5YlUebTxyrT8Yds6IgGz4be+wxN3fvIYHV/rFuTMlbHG75XcNtxGmG6gfnGOY6UKirFiuOjUeyTnHvJwcFM6Jk/OCizjLx0gFiHfzLVZ3d0nXRH/r+KT4u/eeOJ9uJ0dyYZd/3/wLD2zml",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},J=void 0,m={},k=[];function y(e){const n={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Save User Profile"}),"\n",(0,i.jsx)(o(),{method:"put",path:"/user-svc/self",context:"endpoint"}),"\n",(0,i.jsx)(n.p,{children:"Save user's own profile information."}),"\n",(0,i.jsx)(b.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(d(),{parameters:[{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(c(),{title:"Body",body:{content:{"application/json":{schema:{properties:{name:{type:"string"},slug:{type:"string"},thumbnailFileId:{example:"file_fQDxusW8og",type:"string"}},type:"object"}}},description:"Save Profile Request",required:!0}}),"\n",(0,i.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function G(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(y,{...e})}):y(e)}}}]);