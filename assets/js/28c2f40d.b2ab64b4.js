"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7332],{77203:(e,s,t)=>{t.r(s),t.d(s,{assets:()=>x,contentTitle:()=>u,default:()=>v,frontMatter:()=>y,metadata:()=>i,toc:()=>f});const i=JSON.parse('{"id":"1backend/get-permissions-by-role","title":"Get Permissions by Role","description":"Retrieve permissions associated with a specific role ID.","source":"@site/docs/1backend/get-permissions-by-role.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-permissions-by-role","permalink":"/docs/1backend/get-permissions-by-role","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-permissions-by-role","title":"Get Permissions by Role","description":"Retrieve permissions associated with a specific role ID.","sidebar_label":"Get Permissions by Role","hide_title":true,"hide_table_of_contents":true,"api":"eJylVU1v20YQ/SuLObUATcpxAxQ81UYDQ01QG1Z8cnQYL0fkxsvdzexSqkrwvwcjShbtKAUKXyRydz7evHkz7KGiqNmEZLyDEu4osaE1qUDcmhiNd1FhjF4bTFSpjUmNQhUDabMyWrG3pOZ/5pCBD8QoYeYVlFBTuj2GuNreeUuQQUDGlhJxhPKhf518DAYZGHkNmBrIwGFLUIJkmleQAdO3zjBVUCbuKIOoG2oRyh7SNohlTGxcDcOwFOMYvIsU5f7dbCZ/L5PefIQMtHeJXJJbDMEavSuk+BrFpJ+kCCxlJjMGnJAkryZRG08AyQ4HyIxbGI4H/vEr6QTDIGe/nYI3d2u0plJ/LW7+/j9AX3MxJjj/McG9wy41ns2/VL01wfvTFSRih1YtiNfE6gOz57dlGjKIpDs2abuT0RUhE192qYHyYSl9T1iLwuA+EqvFWsMyg5ZS4/fa3ElRzKHoIvFZXOtCFFb0o86GYtpbSSfYR9F2bMWvsF6jbXxM5fvf312cg+Q9wFoI+lEkU3Cvufm8DaS+7E2+gFp5a/2GKvW43U0ZalLoKpX8EzmFelS+WrFvVWpIHcpTn3xtnCJXBW9cyg8j1BBWxMchuty3ekc1HHUZzEcahWncygtO6Q7qXXeoRSMVR7QU/4jG1Z3FxN7l2reT2LdztehC8CzsjiQ1KYWyKM4fUT+Rq8ShkHl4ycLl/MxhMmtSrdHshWujKapgMa08t1KONZpcJMFzyHd9+0mtL/LZi2yxLIrNZpPXrss918XeLxZYB3t2kc/yJrV2N5PEbbxZLcZsR7Bxg3VNnBtf7EwK4ckkKybnV2MhkIHIYYQ/yy/y2RnrXDSQQfAxtegmSK8pqckylObu1+ELHvrjSLxlD+97muifVASLxgmoHUH9XvMPcNC8rNMRSfm8X6fCX2YgAheXvn/ESPdsh0GOv3XEMnzLDNbIBh+FH9noJspzBeUKbaT/KPGXu/0i/1UdF/9J8AeVuq3wjraTN8jgibbHD8OwHLKD3gXIeHmpNYU0cfth1QzTzXD94TNkgPtZPU6HBMsODxL9JKTX0zVCkN8h+4lL34+zNwzP9uPVTz2eR3q0FoaWwzB8B/60qUQ=","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Register","permalink":"/docs/1backend/register"},"next":{"title":"List Roles","permalink":"/docs/1backend/list-roles"}}');var n=t(74848),r=t(28453),o=t(53746),a=t.n(o),c=t(56518),l=t.n(c),p=t(99972),d=t.n(p),b=t(25342),h=t.n(b),m=(t(44215),t(82223),t(24861));const y={id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVU1v20YQ/SuLObUATcpxAxQ81UYDQ01QG1Z8cnQYL0fkxsvdzexSqkrwvwcjShbtKAUKXyRydz7evHkz7KGiqNmEZLyDEu4osaE1qUDcmhiNd1FhjF4bTFSpjUmNQhUDabMyWrG3pOZ/5pCBD8QoYeYVlFBTuj2GuNreeUuQQUDGlhJxhPKhf518DAYZGHkNmBrIwGFLUIJkmleQAdO3zjBVUCbuKIOoG2oRyh7SNohlTGxcDcOwFOMYvIsU5f7dbCZ/L5PefIQMtHeJXJJbDMEavSuk+BrFpJ+kCCxlJjMGnJAkryZRG08AyQ4HyIxbGI4H/vEr6QTDIGe/nYI3d2u0plJ/LW7+/j9AX3MxJjj/McG9wy41ns2/VL01wfvTFSRih1YtiNfE6gOz57dlGjKIpDs2abuT0RUhE192qYHyYSl9T1iLwuA+EqvFWsMyg5ZS4/fa3ElRzKHoIvFZXOtCFFb0o86GYtpbSSfYR9F2bMWvsF6jbXxM5fvf312cg+Q9wFoI+lEkU3Cvufm8DaS+7E2+gFp5a/2GKvW43U0ZalLoKpX8EzmFelS+WrFvVWpIHcpTn3xtnCJXBW9cyg8j1BBWxMchuty3ekc1HHUZzEcahWncygtO6Q7qXXeoRSMVR7QU/4jG1Z3FxN7l2reT2LdztehC8CzsjiQ1KYWyKM4fUT+Rq8ShkHl4ycLl/MxhMmtSrdHshWujKapgMa08t1KONZpcJMFzyHd9+0mtL/LZi2yxLIrNZpPXrss918XeLxZYB3t2kc/yJrV2N5PEbbxZLcZsR7Bxg3VNnBtf7EwK4ckkKybnV2MhkIHIYYQ/yy/y2RnrXDSQQfAxtegmSK8pqckylObu1+ELHvrjSLxlD+97muifVASLxgmoHUH9XvMPcNC8rNMRSfm8X6fCX2YgAheXvn/ESPdsh0GOv3XEMnzLDNbIBh+FH9noJspzBeUKbaT/KPGXu/0i/1UdF/9J8AeVuq3wjraTN8jgibbHD8OwHLKD3gXIeHmpNYU0cfth1QzTzXD94TNkgPtZPU6HBMsODxL9JKTX0zVCkN8h+4lL34+zNwzP9uPVTz2eR3q0FoaWwzB8B/60qUQ=",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},u=void 0,x={},f=[];function k(e){const s={p:"p",...(0,r.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Get Permissions by Role"}),"\n",(0,n.jsx)(a(),{method:"get",path:"/user-svc/role/{roleId}/permissions",context:"endpoint"}),"\n",(0,n.jsx)(s.p,{children:"Retrieve permissions associated with a specific role ID."}),"\n",(0,n.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(l(),{parameters:[{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(d(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(h(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{permissions:{items:{type:"string"},type:"array"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function v(e={}){const{wrapper:s}={...(0,r.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(k,{...e})}):k(e)}}}]);