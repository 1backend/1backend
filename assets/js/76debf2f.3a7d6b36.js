"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1777],{93884:(e,a,t)=>{t.r(a),t.d(a,{assets:()=>f,contentTitle:()=>b,default:()=>m,frontMatter:()=>z,metadata:()=>n,toc:()=>k});const n=JSON.parse('{"id":"1backend/add-user-to-organization","title":"Add a User to an Organization","description":"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.","source":"@site/docs/1backend/add-user-to-organization.api.mdx","sourceDirName":"1backend","slug":"/1backend/add-user-to-organization","permalink":"/docs/1backend/add-user-to-organization","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"add-user-to-organization","title":"Add a User to an Organization","description":"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.","sidebar_label":"Add a User to an Organization","hide_title":true,"hide_table_of_contents":true,"api":"eJzNVk1z2zYQ/SuYPbUzNCnbyUzLU+027ajN1B7LPjk6rIAViQQEGACUonD43ztLijZlqZn04ulFJKH9eG8X+4AWFAXpdR21s5DDlTFuGwRagU0snddfSYkmkBfRCVRKoHWxJP+8JkJNUq+1FM4XaPVX5FCpuC9pMNpqY8SKBIagC0tq6uKdIbHVsdRWxJIOQ0ACribff8wV5IBKPQTy9+5mYgYJ1Oixokg+QP7YvmA0tRXz3yABzcs1xhISsFgR5DDNO1eQgKfPjfakII++oQSCLKlCyFuIu5o9QvTaFtB1y8GYQrx2ascW0tlINvIr1rXRsg+bfQyMp52Eqj3zi5oCf3GxmOZximRccauPJCN0vPSib0qJh31LDhjfDdiOKHW8EGpnw5D9Yjbjx2HUPiIqRUqERkoKYd0Ys4Pk+zkeIe8SeHMq19xu0Ggl/lzc/P1fEhwWkbx3/jtr2CM5P8HaPu/9V0RyeYzkd+dXWimyrwjjzTGM6YbK+j1hXRRr19jXq8/b03smkrdoxIL8hrx418d8HUhdAoFk43Xc9apzTejJXzWxhPxxybIQsWBBGqZosZGwTKCiWDrWstr1Q9nLUA4ZD/9Z2MhsqkRZe6hLXW8GnJjpDmrXeMMBMuMkmtKFmL/96eLyHBjBCHDBhAeOU5gvy3m/q0l82Jt8ALF2fBqQEqtdL9ooSaBVIrpPZAXKQU/E2ruqV++RqHjvCm0FWVU7bWM6am5JqHr8e9W92o/ZqOP7ImOt/6JdX3bu2N2zur77glVtaKqWT03iHGs36i/KvvtUoebyBDQUfgnaFo3B6J1NpasmQG7nYtHUtfPck6GiZYx1nmXnK5SfyCp2yOBYd+dnFqPekKi09I4boyUFURuMa+cr5m60JBt61GO+P27fi81lOjvIFvIs2263aWGb1Pki2/uFDIvanF2ms7SMlWEMkXwVbtaLIdsz2LDFoiCfapf1JhkXVUeuGJxfD0QgAd47A/xZepnOzrxML37muLwpK7RTpHzkP50saMWLo/egGpOj7/92kdhvrUhfYlYb1Jbp9qVv9zP4COMM8r3jkGV+dD/oB3GZAA8c+7btCgM9eNN1vPy5Ic+ysExgg17jilvAVxMd+F1BvkYT6Bv1++Fuf1z/KI5vMCfZjNNj+XzeoGn4CxL4RLvjG0637JJxHhnYYPTrkP6MhWAS5EhDu2T0uJKS6vhN26nq3d4s7iGB1f6mVDnFPh63fD/B7QDY9eXo1apfa8GgLRos2HaIydqAew17Vg2GlIwvzOpkSV6qzkCEf5nWSZe2HTSp657sh7/+1eNJ6gZr7tCy67p/AId++NA=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Create an Organization","permalink":"/docs/1backend/create-organization"},"next":{"title":"Remove a User from an Organization","permalink":"/docs/1backend/remove-user-from-organization"}}');var i=t(74848),o=t(28453),r=t(53746),s=t.n(r),d=t(56518),p=t.n(d),c=t(99972),l=t.n(c),u=t(25342),h=t.n(u),g=(t(44215),t(82223),t(24861));const z={id:"add-user-to-organization",title:"Add a User to an Organization",description:"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.",sidebar_label:"Add a User to an Organization",hide_title:!0,hide_table_of_contents:!0,api:"eJzNVk1z2zYQ/SuYPbUzNCnbyUzLU+027ajN1B7LPjk6rIAViQQEGACUonD43ztLijZlqZn04ulFJKH9eG8X+4AWFAXpdR21s5DDlTFuGwRagU0snddfSYkmkBfRCVRKoHWxJP+8JkJNUq+1FM4XaPVX5FCpuC9pMNpqY8SKBIagC0tq6uKdIbHVsdRWxJIOQ0ACribff8wV5IBKPQTy9+5mYgYJ1Oixokg+QP7YvmA0tRXz3yABzcs1xhISsFgR5DDNO1eQgKfPjfakII++oQSCLKlCyFuIu5o9QvTaFtB1y8GYQrx2ascW0tlINvIr1rXRsg+bfQyMp52Eqj3zi5oCf3GxmOZximRccauPJCN0vPSib0qJh31LDhjfDdiOKHW8EGpnw5D9Yjbjx2HUPiIqRUqERkoKYd0Ys4Pk+zkeIe8SeHMq19xu0Ggl/lzc/P1fEhwWkbx3/jtr2CM5P8HaPu/9V0RyeYzkd+dXWimyrwjjzTGM6YbK+j1hXRRr19jXq8/b03smkrdoxIL8hrx418d8HUhdAoFk43Xc9apzTejJXzWxhPxxybIQsWBBGqZosZGwTKCiWDrWstr1Q9nLUA4ZD/9Z2MhsqkRZe6hLXW8GnJjpDmrXeMMBMuMkmtKFmL/96eLyHBjBCHDBhAeOU5gvy3m/q0l82Jt8ALF2fBqQEqtdL9ooSaBVIrpPZAXKQU/E2ruqV++RqHjvCm0FWVU7bWM6am5JqHr8e9W92o/ZqOP7ImOt/6JdX3bu2N2zur77glVtaKqWT03iHGs36i/KvvtUoebyBDQUfgnaFo3B6J1NpasmQG7nYtHUtfPck6GiZYx1nmXnK5SfyCp2yOBYd+dnFqPekKi09I4boyUFURuMa+cr5m60JBt61GO+P27fi81lOjvIFvIs2263aWGb1Pki2/uFDIvanF2ms7SMlWEMkXwVbtaLIdsz2LDFoiCfapf1JhkXVUeuGJxfD0QgAd47A/xZepnOzrxML37muLwpK7RTpHzkP50saMWLo/egGpOj7/92kdhvrUhfYlYb1Jbp9qVv9zP4COMM8r3jkGV+dD/oB3GZAA8c+7btCgM9eNN1vPy5Ic+ysExgg17jilvAVxMd+F1BvkYT6Bv1++Fuf1z/KI5vMCfZjNNj+XzeoGn4CxL4RLvjG0637JJxHhnYYPTrkP6MhWAS5EhDu2T0uJKS6vhN26nq3d4s7iGB1f6mVDnFPh63fD/B7QDY9eXo1apfa8GgLRos2HaIydqAew17Vg2GlIwvzOpkSV6qzkCEf5nWSZe2HTSp657sh7/+1eNJ6gZr7tCy67p/AId++NA=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,f={},k=[];function j(e){const a={p:"p",...(0,o.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Add a User to an Organization"}),"\n",(0,i.jsx)(s(),{method:"post",path:"/user-svc/organization/{organizationId}/user",context:"endpoint"}),"\n",(0,i.jsx)(a.p,{children:"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization."}),"\n",(0,i.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(p(),{parameters:[{description:"Organization ID",in:"path",name:"organizationId",required:!0,schema:{type:"string"}}]}),"\n",(0,i.jsx)(l(),{title:"Body",body:{content:{"application/json":{schema:{properties:{userId:{type:"string"}},type:"object"}}},description:"Add User to Organization Request",required:!0}}),"\n",(0,i.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"User added successfully",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},403:{description:"Forbidden",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Organization/User not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function m(e={}){const{wrapper:a}={...(0,o.R)(),...e.components};return a?(0,i.jsx)(a,{...e,children:(0,i.jsx)(j,{...e})}):j(e)}}}]);