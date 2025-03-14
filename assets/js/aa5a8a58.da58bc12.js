"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1269],{40127:(s,e,t)=>{t.r(e),t.d(e,{assets:()=>h,contentTitle:()=>x,default:()=>G,frontMatter:()=>m,metadata:()=>n,toc:()=>y});const n=JSON.parse('{"id":"1backend/list-grants","title":"List Grants","description":"List grants.","source":"@site/docs/1backend/list-grants.api.mdx","sourceDirName":"1backend","slug":"/1backend/list-grants","permalink":"/docs/1backend/list-grants","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"list-grants","title":"List Grants","description":"List grants.","sidebar_label":"List Grants","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVsFu2zgQ/RVizrLkJFtgV6d1iqIIGiBB3JwSA6WpscSGIhmSsusVBOxH9Av3SxZDSbEae7tFDwvsxabJmeGb92aGbqFAL5y0QRoNOVxLH1jpuA4+fdSP+n1csgI3UiPbVVJUzKum9Iw7ZNx7WWosmLco5EYKZtHV0ntptE+Y2aJzspC6ZKFCCsIbFZgweiPLxnG6M95yh8+NdOij2afGo5v5rcgjjnwrcfdpEjiFBIzF3v2qgByU9KEHCgk4fG7Qh0tT7CFvQRgdUAdacmuVFNEt++wp3Ra8qLDmtLKOggaJPv56uY5uaCHsLUIOPjipS+gSIBJOHHTJuGPWn1EE6GjrBMcDsXc92gG3dFhAHlyDHW14a7Tv8ZzP5/T1baCbD5D8bIK9xrSSAWt/bCBP5/1DxPhjrMtYNLvKxMKJt2MR9b6dRkwOcI4iDxvcOb4/xfS/W0QxfpmfHcO717wJlXHyDyx+nlR0zrgfrIougTenRL3SAZ3mii3RbdGxdzHmfwOJ5EPROBn2kD+0cIncoVs0oYL8YdWtEgicxH2Ae4+OLbcCVgnUGCpDfWhNrGTLyR6ysY+zcuxNH1PyMXbjFBllygiuKuND/ubX84szoFtGEEtKqs9jCuU1ZR/3FtnjYPIIbGOUMjss2HrPOPOWC2RcFyyYJ9SMi77R2MaZOlbgmAy7NqXUDHVhjdSBBo2k+BXyAkkDzWtibDGUSlQADmVn5Qfs645UuTsMondfeG0VHg+WUY5xnkyqXeqNGQcYF1F3rLkk0jxX6H/3UpeN4sEZnQpTT+DdXrFlY61xpEbPcxWCzbPsbM3FE+qCHDI4mkwLJkxdG80GM7Yxju1N49jiik2qzv/159daCmdIUCnQz9bcYxFJXjdSBRYM84IrJA6VFKh9zH5E+P72mm0v0vk3+HyeZbvdLi11kxpXZoOfz3hp1ewinadVqFXsc3S1v9ks+9sP6fkdL0t0qTRZNMlIHBmIeTi77HOCBKgG+4Tn6UU6nzmRnv8Wh5vxoeZ6gnQyq+EVW5O35f/ybA6FGvBLyKziUlPSUYB26NoHGMNAMr4RqwSoP+mwbUnoe6e6jrafG3Q0KVYJbLmTfE1MP6y6ZGwZavQn3EMOb3uuZtSrJAFXTd8zr0ZZl4weCyHQhu/aTofP7c3yIySwHt792hTk4/iO3la+gxzi/4ZYvmQQ91pQXJcNL8m2j0nty4cxc2hsgpSMC8pqPNL7CcLXg6FPhD4prZMubduPja57se+P/tHjZRr11iTiquu6vwFYgGpG","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Change User Password","permalink":"/docs/1backend/change-password"},"next":{"title":"Save Grants","permalink":"/docs/1backend/save-grants"}}');var i=t(74848),r=t(28453),a=t(53746),c=t.n(a),d=t(56518),o=t.n(d),p=t(99972),l=t.n(p),b=t(25342),g=t.n(b),u=(t(44215),t(82223),t(24861));const m={id:"list-grants",title:"List Grants",description:"List grants.",sidebar_label:"List Grants",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVsFu2zgQ/RVizrLkJFtgV6d1iqIIGiBB3JwSA6WpscSGIhmSsusVBOxH9Av3SxZDSbEae7tFDwvsxabJmeGb92aGbqFAL5y0QRoNOVxLH1jpuA4+fdSP+n1csgI3UiPbVVJUzKum9Iw7ZNx7WWosmLco5EYKZtHV0ntptE+Y2aJzspC6ZKFCCsIbFZgweiPLxnG6M95yh8+NdOij2afGo5v5rcgjjnwrcfdpEjiFBIzF3v2qgByU9KEHCgk4fG7Qh0tT7CFvQRgdUAdacmuVFNEt++wp3Ra8qLDmtLKOggaJPv56uY5uaCHsLUIOPjipS+gSIBJOHHTJuGPWn1EE6GjrBMcDsXc92gG3dFhAHlyDHW14a7Tv8ZzP5/T1baCbD5D8bIK9xrSSAWt/bCBP5/1DxPhjrMtYNLvKxMKJt2MR9b6dRkwOcI4iDxvcOb4/xfS/W0QxfpmfHcO717wJlXHyDyx+nlR0zrgfrIougTenRL3SAZ3mii3RbdGxdzHmfwOJ5EPROBn2kD+0cIncoVs0oYL8YdWtEgicxH2Ae4+OLbcCVgnUGCpDfWhNrGTLyR6ysY+zcuxNH1PyMXbjFBllygiuKuND/ubX84szoFtGEEtKqs9jCuU1ZR/3FtnjYPIIbGOUMjss2HrPOPOWC2RcFyyYJ9SMi77R2MaZOlbgmAy7NqXUDHVhjdSBBo2k+BXyAkkDzWtibDGUSlQADmVn5Qfs645UuTsMondfeG0VHg+WUY5xnkyqXeqNGQcYF1F3rLkk0jxX6H/3UpeN4sEZnQpTT+DdXrFlY61xpEbPcxWCzbPsbM3FE+qCHDI4mkwLJkxdG80GM7Yxju1N49jiik2qzv/159daCmdIUCnQz9bcYxFJXjdSBRYM84IrJA6VFKh9zH5E+P72mm0v0vk3+HyeZbvdLi11kxpXZoOfz3hp1ewinadVqFXsc3S1v9ks+9sP6fkdL0t0qTRZNMlIHBmIeTi77HOCBKgG+4Tn6UU6nzmRnv8Wh5vxoeZ6gnQyq+EVW5O35f/ybA6FGvBLyKziUlPSUYB26NoHGMNAMr4RqwSoP+mwbUnoe6e6jrafG3Q0KVYJbLmTfE1MP6y6ZGwZavQn3EMOb3uuZtSrJAFXTd8zr0ZZl4weCyHQhu/aTofP7c3yIySwHt792hTk4/iO3la+gxzi/4ZYvmQQ91pQXJcNL8m2j0nty4cxc2hsgpSMC8pqPNL7CcLXg6FPhD4prZMubduPja57se+P/tHjZRr11iTiquu6vwFYgGpG",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},x=void 0,h={},y=[];function j(s){const e={code:"code",p:"p",...(0,r.R)(),...s.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(u.default,{as:"h1",className:"openapi__heading",children:"List Grants"}),"\n",(0,i.jsx)(c(),{method:"post",path:"/user-svc/grants",context:"endpoint"}),"\n",(0,i.jsx)(e.p,{children:"List grants."}),"\n",(0,i.jsx)(e.p,{children:"Grants define which slugs are assigned specific permissions, overriding the default configuration."}),"\n",(0,i.jsxs)(e.p,{children:["Requires the ",(0,i.jsx)(e.code,{children:"user-svc:grant:view"})," permission."]}),"\n",(0,i.jsx)(u.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(o(),{parameters:void 0}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{permissionId:{type:"string"},slug:{type:"string"}},type:"object"}}},description:"List Grants Request",required:!0}}),"\n",(0,i.jsx)(g(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{grants:{items:{properties:{id:{type:"string"},permissionId:{type:"string"},slugs:{description:"Slugs who are granted the PermissionId",items:{type:"string"},type:"array"}},type:"object"},type:"array"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function G(s={}){const{wrapper:e}={...(0,r.R)(),...s.components};return e?(0,i.jsx)(e,{...s,children:(0,i.jsx)(j,{...s})}):j(s)}}}]);