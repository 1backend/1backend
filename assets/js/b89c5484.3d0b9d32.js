"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[2834],{83594:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>v,contentTitle:()=>u,default:()=>j,frontMatter:()=>g,metadata:()=>s,toc:()=>x});const s=JSON.parse('{"id":"1backend/remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","source":"@site/docs/1backend/remove-instance.api.mdx","sourceDirName":"1backend","slug":"/1backend/remove-instance","permalink":"/docs/1backend/remove-instance","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"remove-instance","title":"Remove Instance","description":"Removes a registered instance by ID.","sidebar_label":"Remove Instance","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VU1v2zgQ/SvEnHYBRbLrLhbQadNNUBgNukXSnFIfaGossaFIdkjZ6xX43xcjyYmduIs95WLr443mveF7ZA8VBkXaR+0slHCLrdtiEFIQ1jpEJKyEtiFKq1Cs92J5lUMGziNJLllWUAINRcsJBRl4SbLFiBSgfOhftDjgxPIKMtD8yMvYQAZWtggl6AoyIPzRacIKykgdZhBUg62Esoe494wKkbStIaUVg4N3NmDg9+9m7/nvtOlnJ/50NqKNkDJ4P5u9hiztVhpdjbTUBC57kN4brQa1xffA0P6IjSeeRdRjbyRydI5kdnji1t9RRUgpDTzmr3ncW9nFxpH+B6s3ZHJmaHdIW61QWBfFxnX27ej8dn6BIpKVRjAvJHE9fPNtKKUMAqqOdNwPlv6AkpAuu9hA+bBiD0ZZs9vhdsgN7cXdVsEqgxZj4zglFRqMYzq4CgqakBdhq4pDxopeVwm4G2sc89ORYXxhnJKmcSGW8/li8Ttw2wOrO1Y5Cjvm9nKGX/cexbcJ8g3ExhnjdlhxtKUIXioU0lYiuke0QqoxhGJDrhWxQXEfkFiZuHG1tgJt5Z22MT8kuUFZIT1n+XKy8rAk8DRZ6fUn3A+z1nbjmCevolTDKmIrNSsO0mD4I2hbd0ZGcjZXrj369peluOu8dxQhm4bUxOjLopivpXpEW3FBwYk/ncLl8sLKqLcoWq3IhdHnQXgj48ZRy3KMVmgDMp9Dv49fbsR2kc9OuoWyKHa7XV7bLndUF1NdKGTtzcUin+VNbA1ziEht+GszpeqZbNjJukbKtSsGSMFz0tEwZP5hFAIZsB1G+rN8kc8uSOXsgQy8C7GV9ojpuIuLox35RH//HJn/u+FP6xbx71h4I7XlxsMQ+snPD3Ds58EPT91LXXEU2LoM7Pu1DHhPJiV+/KND4lStMthK0nLNyvnY0IGvKyg30gT8DxG/3E6nxa/i9HQ5S/vgQbvnqUrT8R1k8Ij78fRJq5QdnMxExheXSqGPRyWvNpt0HPer65vrr9eQgZyC+Gx9/l52uOAGZxm9jM7Ign9T9pOSvh+DldITfnz104qnvE7iWURK6V+qxcMF","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register Instance","permalink":"/docs/1backend/register-instance"},"next":{"title":"List Service Instances","permalink":"/docs/1backend/list-instances"}}');var i=t(74848),r=t(28453),a=t(53746),c=t.n(a),o=t(56518),d=t.n(o),p=t(99972),l=t.n(p),h=t(25342),b=t.n(h),m=(t(44215),t(82223),t(24861));const g={id:"remove-instance",title:"Remove Instance",description:"Removes a registered instance by ID.",sidebar_label:"Remove Instance",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VU1v2zgQ/SvEnHYBRbLrLhbQadNNUBgNukXSnFIfaGossaFIdkjZ6xX43xcjyYmduIs95WLr443mveF7ZA8VBkXaR+0slHCLrdtiEFIQ1jpEJKyEtiFKq1Cs92J5lUMGziNJLllWUAINRcsJBRl4SbLFiBSgfOhftDjgxPIKMtD8yMvYQAZWtggl6AoyIPzRacIKykgdZhBUg62Esoe494wKkbStIaUVg4N3NmDg9+9m7/nvtOlnJ/50NqKNkDJ4P5u9hiztVhpdjbTUBC57kN4brQa1xffA0P6IjSeeRdRjbyRydI5kdnji1t9RRUgpDTzmr3ncW9nFxpH+B6s3ZHJmaHdIW61QWBfFxnX27ej8dn6BIpKVRjAvJHE9fPNtKKUMAqqOdNwPlv6AkpAuu9hA+bBiD0ZZs9vhdsgN7cXdVsEqgxZj4zglFRqMYzq4CgqakBdhq4pDxopeVwm4G2sc89ORYXxhnJKmcSGW8/li8Ttw2wOrO1Y5Cjvm9nKGX/cexbcJ8g3ExhnjdlhxtKUIXioU0lYiuke0QqoxhGJDrhWxQXEfkFiZuHG1tgJt5Z22MT8kuUFZIT1n+XKy8rAk8DRZ6fUn3A+z1nbjmCevolTDKmIrNSsO0mD4I2hbd0ZGcjZXrj369peluOu8dxQhm4bUxOjLopivpXpEW3FBwYk/ncLl8sLKqLcoWq3IhdHnQXgj48ZRy3KMVmgDMp9Dv49fbsR2kc9OuoWyKHa7XV7bLndUF1NdKGTtzcUin+VNbA1ziEht+GszpeqZbNjJukbKtSsGSMFz0tEwZP5hFAIZsB1G+rN8kc8uSOXsgQy8C7GV9ojpuIuLox35RH//HJn/u+FP6xbx71h4I7XlxsMQ+snPD3Ds58EPT91LXXEU2LoM7Pu1DHhPJiV+/KND4lStMthK0nLNyvnY0IGvKyg30gT8DxG/3E6nxa/i9HQ5S/vgQbvnqUrT8R1k8Ij78fRJq5QdnMxExheXSqGPRyWvNpt0HPer65vrr9eQgZyC+Gx9/l52uOAGZxm9jM7Ign9T9pOSvh+DldITfnz104qnvE7iWURK6V+qxcMF",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},u=void 0,v={},x=[];function I(e){const n={p:"p",...(0,r.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Remove Instance"}),"\n",(0,i.jsx)(c(),{method:"delete",path:"/registry-svc/instance/{id}",context:"endpoint"}),"\n",(0,i.jsx)(n.p,{children:"Removes a registered instance by ID."}),"\n",(0,i.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(d(),{parameters:[{description:"Instance ID",in:"path",name:"id",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,i.jsx)(b(),{id:void 0,label:void 0,responses:{204:{description:"No Content"},400:{description:"Invalid ID",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Service not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function j(e={}){const{wrapper:n}={...(0,r.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(I,{...e})}):I(e)}}}]);