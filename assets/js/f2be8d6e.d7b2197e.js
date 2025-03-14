"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[4543],{88448:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>K,contentTitle:()=>y,default:()=>j,frontMatter:()=>h,metadata:()=>r,toc:()=>b});const r=JSON.parse('{"id":"openorch/start-default-model","title":"Start the Default Model","description":"Starts The Default Model.","source":"@site/docs/openorch/start-default-model.api.mdx","sourceDirName":"openorch","slug":"/openorch/start-default-model","permalink":"/docs/openorch/start-default-model","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"start-default-model","title":"Start the Default Model","description":"Starts The Default Model.","sidebar_label":"Start the Default Model","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE1v2zgQ/SvEnGXJSVpgV6em2MUi/XJQJyfHQMfUWGJDkSxJ2esK+u/FULbjuOlhgUUvtkAO37x5fI89VBSkVy4qa6CEeUQfg7hrSPxFa+x0FB9tRTp/MA/mM33rlKcgYkPiS8vrk7CRZfoqpSeM9EU48q0KQVmTQwbWkUcGv6mghMDwe+CECxl4Cs6aQAHKHi6nU/57Tmr2HjKQ1kQykXfROa1kQi2+Bi7pIciGWuSvuHMEJdjVV5IRhmEYMnj1EuyN2aBWlXg3n336Lw2c56GiGhmT99af9A3RK1MDd32ZycXPTO4NdrGxXn2n6rcxef2yJpG8QS3m5Dfkxd8J8/dQGjIIJDuv4g7KRQ9vCT356y42UC6WwzKDiHWAcgHJOmK+kbDMoKXYWDaX6yJk4JDroTjas6hGv03SSpEsCNyKBwypU+c1Hym0lagbG2L5+o/LqwvgngdKcx5xnOqU2LmAdztH4mFf8gBibbW2W6rEaidQBIeSBJpKRPtIRqBMiarE2ts2xeo+kOfJxAdbKyPIVM4qEzlKivEbwor4Rgy2rN/13jjpPuAoKzr1nnZJaGXWlnnyFaJMV0gtKp44oKbwJihTdxqjtyaXtj3Bvr0R8845mwQbRWpidGVRWEfGetnk1tcFDNmZCtdCo6k7rGmCtbEhKilaJb1l1ZWkINYeW9pa/yjW1otVp3SlTC2ub8SJvwKPrZUkE4h5H3j9c/tBbK7y6TNWoSyK7Xab16ZLrPbnQoG105OrfJo3sdXMNZJvw2w9H7k8DRW2WNfkc2WLVFKwnipqLpk5MjMvG8iAbTOOOc2v8unEy/zyT8Z1NsQWzQnT9Jyme332msKZXv1Tvv7/F3jviEj/xsJpVIapJtn6fVgWcERKzE7iwkFJgVlmwMHg4r5fYaB7r4eBl7915Dmwyww26BWuWK/FcsgOXuWEPdKOL05KcuylDeputOnZWzKcBvr2/g4ywH3OnpzNYNnhg9EPW2Z3gn2ejJEC/w7ZL470/ZibYTjWj1u/PHGM41jN4i6HYfgBiDOIBg==","sidebar_class_name":"put api-method","info_path":"docs/1backend/1backend","custom_edit_url":null},"sidebar":"tutorialSidebar","previous":{"title":"Set Role Permissions","permalink":"/docs/openorch/set-role-permission"},"next":{"title":"Start a Model","permalink":"/docs/openorch/start-model"}}');var a=n(74848),o=n(28453),i=n(53746),s=n.n(i),d=n(56518),l=n.n(d),c=n(99972),p=n.n(c),u=n(25342),m=n.n(u),f=(n(44215),n(82223),n(24861));const h={id:"start-default-model",title:"Start the Default Model",description:"Starts The Default Model.",sidebar_label:"Start the Default Model",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE1v2zgQ/SvEnGXJSVpgV6em2MUi/XJQJyfHQMfUWGJDkSxJ2esK+u/FULbjuOlhgUUvtkAO37x5fI89VBSkVy4qa6CEeUQfg7hrSPxFa+x0FB9tRTp/MA/mM33rlKcgYkPiS8vrk7CRZfoqpSeM9EU48q0KQVmTQwbWkUcGv6mghMDwe+CECxl4Cs6aQAHKHi6nU/57Tmr2HjKQ1kQykXfROa1kQi2+Bi7pIciGWuSvuHMEJdjVV5IRhmEYMnj1EuyN2aBWlXg3n336Lw2c56GiGhmT99af9A3RK1MDd32ZycXPTO4NdrGxXn2n6rcxef2yJpG8QS3m5Dfkxd8J8/dQGjIIJDuv4g7KRQ9vCT356y42UC6WwzKDiHWAcgHJOmK+kbDMoKXYWDaX6yJk4JDroTjas6hGv03SSpEsCNyKBwypU+c1Hym0lagbG2L5+o/LqwvgngdKcx5xnOqU2LmAdztH4mFf8gBibbW2W6rEaidQBIeSBJpKRPtIRqBMiarE2ts2xeo+kOfJxAdbKyPIVM4qEzlKivEbwor4Rgy2rN/13jjpPuAoKzr1nnZJaGXWlnnyFaJMV0gtKp44oKbwJihTdxqjtyaXtj3Bvr0R8845mwQbRWpidGVRWEfGetnk1tcFDNmZCtdCo6k7rGmCtbEhKilaJb1l1ZWkINYeW9pa/yjW1otVp3SlTC2ub8SJvwKPrZUkE4h5H3j9c/tBbK7y6TNWoSyK7Xab16ZLrPbnQoG105OrfJo3sdXMNZJvw2w9H7k8DRW2WNfkc2WLVFKwnipqLpk5MjMvG8iAbTOOOc2v8unEy/zyT8Z1NsQWzQnT9Jyme332msKZXv1Tvv7/F3jviEj/xsJpVIapJtn6fVgWcERKzE7iwkFJgVlmwMHg4r5fYaB7r4eBl7915Dmwyww26BWuWK/FcsgOXuWEPdKOL05KcuylDeputOnZWzKcBvr2/g4ywH3OnpzNYNnhg9EPW2Z3gn2ejJEC/w7ZL470/ZibYTjWj1u/PHGM41jN4i6HYfgBiDOIBg==",sidebar_class_name:"put api-method",info_path:"docs/1backend/1backend",custom_edit_url:null},y=void 0,K={},b=[];function g(e){const t={code:"code",p:"p",...(0,o.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(f.default,{as:"h1",className:"openapi__heading",children:"Start the Default Model"}),"\n",(0,a.jsx)(s(),{method:"put",path:"/model-svc/default-model/start",context:"endpoint"}),"\n",(0,a.jsx)(t.p,{children:"Starts The Default Model."}),"\n",(0,a.jsxs)(t.p,{children:["Requires the ",(0,a.jsx)(t.code,{children:"model-svc:model:create"})," permission."]}),"\n",(0,a.jsx)(l(),{parameters:void 0}),"\n",(0,a.jsx)(p(),{title:"Body",body:void 0}),"\n",(0,a.jsx)(m(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function j(e={}){const{wrapper:t}={...(0,o.R)(),...e.components};return t?(0,a.jsx)(t,{...e,children:(0,a.jsx)(g,{...e})}):g(e)}}}]);