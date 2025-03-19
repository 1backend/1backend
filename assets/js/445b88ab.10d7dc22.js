"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1663],{15974:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>j,contentTitle:()=>k,default:()=>w,frontMatter:()=>y,metadata:()=>s,toc:()=>m});const s=JSON.parse('{"id":"1backend/checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","source":"@site/docs/1backend/checkout-repo.api.mdx","sourceDirName":"1backend","slug":"/1backend/checkout-repo","permalink":"/docs/1backend/checkout-repo","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","sidebar_label":"Checkout a git repository","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVt9v2zYQ/lcIPnWAIjnJCmx6mtN1S7ZiMaLkKTEKmjpLbCiSJSmrmqH/fThKspXYCYI+FNiLIR+P3/3+jluag+NWGC+0oin9UAJ/1LUnjBTCEwtGO+G1bYnegCWl98YRbYlzJWGo5QxwsRacbMA6oRURymvCiIfKaMtsS3JhgSNE/KAWYNfaVg4vlkxK3RAutQLSCF+SSihRMUlK4YLJtbZkzZwHS/jgV0wjqg1Yhv5e5TSl48kNGE0jauFrDc5f6Lyl6ZZyrTwoj5/MGCl4uJh8cRjtljpeQsXwy1iE9QJc+Meca7TN8ftpghbDCSbB60dQwcvL29tFRljtSxpR3xqgKXXeClXQLqLOlZ8foT0Ey7JLYqzYMA/kEVryTocTJgMonnKtFHAU/vQK8mfTvObqCDY1JdYEFLet8ZDv7R41EsI8hL89Ev1hBEcRaysP8f6opZw23N3NJ/IO4iKO+rZLk6QQvqxXMddVUjuwCWofN+DAKlbBoZW74WTi+JAdvPNG/4deP0S/sEzxMiKeFRHicl1VwpPscn4I0+0kevUFuKcdil4YR+xuctP39tDlwkJOU29r6FDgjFau796z2exIr9Wcg3PrWsq2HyfICUL7EiZpp9H3zkwu7KHV38fhJ00JFp4ZIw1zU1/elqQuoj8fi/BKbZgUOfkru/7n+8MAa3UI5K2enB5pMoXDoK34F/If5sn74znx2O+SZGCRwD8GzB/jEvIT8NoK39L0fksvgFmwcyTJ9H7ZLSPqWeFoek8zXVsOJNtwuoxoBb7USO1Gh2Y3DG/QxAWtE7fhYfCTkfkp2sHoXDATyIUmidScyVI7n77/5ez8lKLB0Z8M4+tDmnp1QHGtAfIwqDxQsta4siAnqzasPsaBMJUPi4DxfijJ2uoqdDqSDUZFPulCKAIqN1qosMQE4pfAcsBy9FxF50PXhGLsh4EZ8Te0IctYoJv9ivv4jVVGwtOVNVZlsngORP3G2IsHjt8L+izu/+4IdS/bseCEGYVa63HvMh76CyomAhaT4H5zQhW1ZN5qhTw+iX1xRbLaGG39zjzyfpokpyvGH0HlgfjpAUnOSa3EWmBZej2SwwakNhUoT4xkHt8cgdArwa3GZhEc3MmKOcjJ/IpMRsBhcaTgoFxI6+jdn4tPZHMez574hjupaZq4UHWsbZEM91zCCiNPzuNZXPpKhh0KtnLX66w3vQ/NNawowMZCJ0ElwVoIjyWlpxd9OE9SPYvP49mJ5fHZr4iLI1IxNfH0xRccfZa3yePof/HsG4bBwzefGMmEmrwkeoq4p3uKCEsyvAh3NLGMKNIB6m23WPs7K7sOxV9rsMhRy4humBVshfm/X3bROKHIK/0gfehzdoLUgIVhsu5H9BmJdtF4Y845GP+q7pT0FtfZLY3oanjAVjrHO5Y1GBFraErDEzj0KioE2ZZKpoqaFajbYyJbsIHV9jyCLkXjB0Y1Hql24uFzHuoDwV8M6+iV7bZnqa7b6fdHL97YkV+vjfVcdl33H5cRYI8=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Save Secrets","permalink":"/docs/1backend/save-secrets"},"next":{"title":"Reset Password","permalink":"/docs/1backend/reset-password"}}');var i=o(74848),r=o(28453),a=o(53746),n=o.n(a),c=o(56518),p=o.n(c),d=o(99972),h=o.n(d),l=o(25342),u=o.n(l),f=(o(44215),o(82223),o(24861));const y={id:"checkout-repo",title:"Checkout a git repository",description:"Checkout a git repository over https or ssh at a specific version into a temporary directory.",sidebar_label:"Checkout a git repository",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVt9v2zYQ/lcIPnWAIjnJCmx6mtN1S7ZiMaLkKTEKmjpLbCiSJSmrmqH/fThKspXYCYI+FNiLIR+P3/3+jluag+NWGC+0oin9UAJ/1LUnjBTCEwtGO+G1bYnegCWl98YRbYlzJWGo5QxwsRacbMA6oRURymvCiIfKaMtsS3JhgSNE/KAWYNfaVg4vlkxK3RAutQLSCF+SSihRMUlK4YLJtbZkzZwHS/jgV0wjqg1Yhv5e5TSl48kNGE0jauFrDc5f6Lyl6ZZyrTwoj5/MGCl4uJh8cRjtljpeQsXwy1iE9QJc+Meca7TN8ftpghbDCSbB60dQwcvL29tFRljtSxpR3xqgKXXeClXQLqLOlZ8foT0Ey7JLYqzYMA/kEVryTocTJgMonnKtFHAU/vQK8mfTvObqCDY1JdYEFLet8ZDv7R41EsI8hL89Ev1hBEcRaysP8f6opZw23N3NJ/IO4iKO+rZLk6QQvqxXMddVUjuwCWofN+DAKlbBoZW74WTi+JAdvPNG/4deP0S/sEzxMiKeFRHicl1VwpPscn4I0+0kevUFuKcdil4YR+xuctP39tDlwkJOU29r6FDgjFau796z2exIr9Wcg3PrWsq2HyfICUL7EiZpp9H3zkwu7KHV38fhJ00JFp4ZIw1zU1/elqQuoj8fi/BKbZgUOfkru/7n+8MAa3UI5K2enB5pMoXDoK34F/If5sn74znx2O+SZGCRwD8GzB/jEvIT8NoK39L0fksvgFmwcyTJ9H7ZLSPqWeFoek8zXVsOJNtwuoxoBb7USO1Gh2Y3DG/QxAWtE7fhYfCTkfkp2sHoXDATyIUmidScyVI7n77/5ez8lKLB0Z8M4+tDmnp1QHGtAfIwqDxQsta4siAnqzasPsaBMJUPi4DxfijJ2uoqdDqSDUZFPulCKAIqN1qosMQE4pfAcsBy9FxF50PXhGLsh4EZ8Te0IctYoJv9ivv4jVVGwtOVNVZlsngORP3G2IsHjt8L+izu/+4IdS/bseCEGYVa63HvMh76CyomAhaT4H5zQhW1ZN5qhTw+iX1xRbLaGG39zjzyfpokpyvGH0HlgfjpAUnOSa3EWmBZej2SwwakNhUoT4xkHt8cgdArwa3GZhEc3MmKOcjJ/IpMRsBhcaTgoFxI6+jdn4tPZHMez574hjupaZq4UHWsbZEM91zCCiNPzuNZXPpKhh0KtnLX66w3vQ/NNawowMZCJ0ElwVoIjyWlpxd9OE9SPYvP49mJ5fHZr4iLI1IxNfH0xRccfZa3yePof/HsG4bBwzefGMmEmrwkeoq4p3uKCEsyvAh3NLGMKNIB6m23WPs7K7sOxV9rsMhRy4humBVshfm/X3bROKHIK/0gfehzdoLUgIVhsu5H9BmJdtF4Y845GP+q7pT0FtfZLY3oanjAVjrHO5Y1GBFraErDEzj0KioE2ZZKpoqaFajbYyJbsIHV9jyCLkXjB0Y1Hql24uFzHuoDwV8M6+iV7bZnqa7b6fdHL97YkV+vjfVcdl33H5cRYI8=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},k=void 0,j={},m=[];function b(e){const t={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Checkout a git repository"}),"\n",(0,i.jsx)(n(),{method:"post",path:"/source-svc/repo/checkout",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Checkout a git repository over https or ssh at a specific version into a temporary directory.\nPerforms a shallow clone with minimal history for faster checkout."}),"\n",(0,i.jsx)(f.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(p(),{parameters:void 0}),"\n",(0,i.jsx)(h(),{title:"Body",body:{content:{"application/json":{schema:{properties:{password:{description:"Password or token for HTTPS auth",type:"string"},ssh_key:{description:"SSH private key (optional for SSH connection)",type:"string"},ssh_key_pwd:{description:"Password for SSH private key if encrypted (optional)",type:"string"},token:{description:"Token for HTTPS auth (optional for SSH)",type:"string"},url:{description:"Full repository URL (e.g., https://github.com/user/repo)",type:"string"},username:{description:"Username for HTTPS or SSH user (optional for SSH)",type:"string"},version:{description:"Branch, tag, or commit SHA",type:"string"}},type:"object"}}},description:"Checkout Repo Request",required:!0}}),"\n",(0,i.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Successfully checked out the repository",content:{"application/json":{schema:{properties:{dir:{description:"Directory where the repository was checked out",type:"string"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function w(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(b,{...e})}):b(e)}}}]);