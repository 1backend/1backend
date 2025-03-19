"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[8842],{53921:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>f,contentTitle:()=>m,default:()=>b,frontMatter:()=>g,metadata:()=>i,toc:()=>k});const i=JSON.parse('{"id":"1backend/save-config","title":"Save Config","description":"Save the provided configuration to the server","source":"@site/docs/1backend/save-config.api.mdx","sourceDirName":"1backend","slug":"/1backend/save-config","permalink":"/docs/1backend/save-config","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"save-config","title":"Save Config","description":"Save the provided configuration to the server","sidebar_label":"Save Config","hide_title":true,"hide_table_of_contents":true,"api":"eJylVU1v4zYQ/SvEnBXJSbpAq1OTYlGkXaDBenPy+jCmxjI3FMklKXtdQf+9GFJOlDhd9ONi0OR8vHkzbzRAQ0F65aKyBmpY4p5E3JFw3u5VQ42Q1mxV23tkCxFteg3k9+ShAOsov9w1UEPAPf2S7KEAT197CvHWNkeoB5DWRDKRj+icVjK5VV8C5x0gyB11yCfnOWhUFCY3Dnd232BM1tg0igOhvp+9R99TAfHoCGqwmy8kI4xFcvptSjg9huiVaflRNW9eG+woOJT0xuuYq1SeGqhXGdP6LO94fsNYznnP1ImPmTiYB+eCUrbgrAmZgqvFIjHxnUDZGop/zv4Z0rGAHxaX54keDPZxZ736k5r/kOCJQk7w7q1K7kwkb1CLZRo28d576/9fprGAQLL3Kh6hXg1wS+jJ3/RxB/VqPXLrsA3cyonB5V5yQzuKO8sD7npui0N2gCrP5kXYy+kIHJ/BhhS+95rNKm0l6p0NsX7349X1JXCiE44lw839nKN5TcanoyPxeTL5DGJrtbYHasTmKFCk+RRoGhHtIxmBMs+N2HrbJcE+BPJcjfhgW2UEmcZZZWIJBSiOvyNskqB53KGGm6m3iVt4GmB06nc6ppFmvj8+S/z9N+ycppeSnUT6QnmnfmTBPf+b6WyuL2W29rQ+UKauU4dKp22jKfwclGl7jdFbU0rbzUq4vxPL3jnruWe5F7sYXV1VlxuUj2QadqjgTIw3ojdqq5jdbCca2pO2riMThdMYt9Z3Ymu96JT0lnuuJIWLDQZqxM2dmA1mYI61ksRCrIcndL/efxD763LxAluoq+pwOJSt6Uvr22ryCxW2Tl9cl4tyFzvNiCP5LvyxXebUz6WFA7Yt+VLZKplU3DwVuTNweZvLgQJ4RnOxi/K6XFx4WV79xHGdDbFDM0M62yjwiqnZVv/XX45poiJ9i5XTqAxnT0wMk8BW8CywrHuGsC6ApcTPw8CEP3g9jnz9tSfPul4XsEevcMNFr9ZjcZpu1uQjHaFmeTPsC5YVs4G6z+P9ap+MxcnjRkpy8bu280Vx//AJCthM377ONuzi8cBLHQ9QQ/p2pvlgg3Q3gEbT9tiybQ7JCsBpITxLkBEVpwMXdXoyxxnA1xLOdfAvV/WmyzBkgY/jk31++luPp72RrbmL63Ec/wLWfN58","sidebar_class_name":"put api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Get Config","permalink":"/docs/1backend/get-config"},"next":{"title":"Run a Container","permalink":"/docs/1backend/run-container"}}');var a=t(74848),o=t(28453),s=t(53746),r=t.n(s),c=t(56518),p=t.n(c),d=t(99972),l=t.n(d),u=t(25342),v=t.n(u),x=(t(44215),t(82223),t(24861));const g={id:"save-config",title:"Save Config",description:"Save the provided configuration to the server",sidebar_label:"Save Config",hide_title:!0,hide_table_of_contents:!0,api:"eJylVU1v4zYQ/SvEnBXJSbpAq1OTYlGkXaDBenPy+jCmxjI3FMklKXtdQf+9GFJOlDhd9ONi0OR8vHkzbzRAQ0F65aKyBmpY4p5E3JFw3u5VQ42Q1mxV23tkCxFteg3k9+ShAOsov9w1UEPAPf2S7KEAT197CvHWNkeoB5DWRDKRj+icVjK5VV8C5x0gyB11yCfnOWhUFCY3Dnd232BM1tg0igOhvp+9R99TAfHoCGqwmy8kI4xFcvptSjg9huiVaflRNW9eG+woOJT0xuuYq1SeGqhXGdP6LO94fsNYznnP1ImPmTiYB+eCUrbgrAmZgqvFIjHxnUDZGop/zv4Z0rGAHxaX54keDPZxZ736k5r/kOCJQk7w7q1K7kwkb1CLZRo28d576/9fprGAQLL3Kh6hXg1wS+jJ3/RxB/VqPXLrsA3cyonB5V5yQzuKO8sD7npui0N2gCrP5kXYy+kIHJ/BhhS+95rNKm0l6p0NsX7349X1JXCiE44lw839nKN5TcanoyPxeTL5DGJrtbYHasTmKFCk+RRoGhHtIxmBMs+N2HrbJcE+BPJcjfhgW2UEmcZZZWIJBSiOvyNskqB53KGGm6m3iVt4GmB06nc6ppFmvj8+S/z9N+ycppeSnUT6QnmnfmTBPf+b6WyuL2W29rQ+UKauU4dKp22jKfwclGl7jdFbU0rbzUq4vxPL3jnruWe5F7sYXV1VlxuUj2QadqjgTIw3ojdqq5jdbCca2pO2riMThdMYt9Z3Ymu96JT0lnuuJIWLDQZqxM2dmA1mYI61ksRCrIcndL/efxD763LxAluoq+pwOJSt6Uvr22ryCxW2Tl9cl4tyFzvNiCP5LvyxXebUz6WFA7Yt+VLZKplU3DwVuTNweZvLgQJ4RnOxi/K6XFx4WV79xHGdDbFDM0M62yjwiqnZVv/XX45poiJ9i5XTqAxnT0wMk8BW8CywrHuGsC6ApcTPw8CEP3g9jnz9tSfPul4XsEevcMNFr9ZjcZpu1uQjHaFmeTPsC5YVs4G6z+P9ap+MxcnjRkpy8bu280Vx//AJCthM377ONuzi8cBLHQ9QQ/p2pvlgg3Q3gEbT9tiybQ7JCsBpITxLkBEVpwMXdXoyxxnA1xLOdfAvV/WmyzBkgY/jk31++luPp72RrbmL63Ec/wLWfN58",sidebar_class_name:"put api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},m=void 0,f={},k=[];function h(e){const n={p:"p",...(0,o.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(x.default,{as:"h1",className:"openapi__heading",children:"Save Config"}),"\n",(0,a.jsx)(r(),{method:"put",path:"/config-svc/config",context:"endpoint"}),"\n",(0,a.jsx)(n.p,{children:"Save the provided configuration to the server"}),"\n",(0,a.jsx)(x.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(p(),{parameters:void 0}),"\n",(0,a.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{config:{properties:{data:{additionalProperties:!0,type:"object"},dataJson:{type:"string"},id:{type:"string"},namespace:{type:"string"}},required:["data"],type:"object"}},type:"object"}}},description:"Save Config Request",required:!0}}),"\n",(0,a.jsx)(v(),{id:void 0,label:void 0,responses:{200:{description:"Save Config Response",content:{"application/json":{schema:{type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function b(e={}){const{wrapper:n}={...(0,o.R)(),...e.components};return n?(0,a.jsx)(n,{...e,children:(0,a.jsx)(h,{...e})}):h(e)}}}]);