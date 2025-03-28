"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[3377],{29474:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>z,contentTitle:()=>h,default:()=>j,frontMatter:()=>b,metadata:()=>t,toc:()=>S});const t=JSON.parse('{"id":"1backend/remove-user-from-organization","title":"Remove a User from an Organization","description":"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.","source":"@site/docs/1backend/remove-user-from-organization.api.mdx","sourceDirName":"1backend","slug":"/1backend/remove-user-from-organization","permalink":"/docs/1backend/remove-user-from-organization","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"remove-user-from-organization","title":"Remove a User from an Organization","description":"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.","sidebar_label":"Remove a User from an Organization","hide_title":true,"hide_table_of_contents":true,"api":"eJzNVk1v2zgQ/SvEnHYBRXLqFljotMk2XXg32BZxekp9GFNjiQ1FqiRl1xX03xcjybH8kQLtAsFebImajzdvyMdpICMvnaqCsgZSuNLabrxAI7AOhXXqG2Wi9uREsAKzTKCxoSC3XxO+IqlWSgrrcjTqG3KoWNwX1BttlNZiSQK9V7mhbOzirCaxUaFQRoSCDkNABLYi173MMkjBUWnX9NGTe+ds+X5kCxFU6LCkQM5D+tAclTW2FbO3EIHi5QpDAREYLAlSGCefZRCBoy+1cpRBGlxNEXhZUImQNhC2FXv44JTJoW2j44QM8tlETMsPJlj0xuTDtc22bCGtCWQCP2JVaSU73Mlnz/mb01B2+ZlkgLZtoyOsdx2tooPMxIoDtu76tNAyAl9Z48lz2FeTCf+dKRuzjDLhaynJ+1Wt9Rai/wC3jeD1uVwzs0atMvHX/P0/P5KgcryrgurLIOesO9vSZ5Bcnqna7A/LCyKZniJ5Z91SZRmZF4Tx+hTGeAMl3Z4wNoiVrc3L8fPm/J4J5AxqMSe3JiduupgvA6mNwJOsnQrbTqGuCR25qzoUkD4s+IQHzFm8+lM0X0tYRFBSKCyLX0aaAnVCxx6QsIpc+LVMxrqVNIcq1nZmSdNLTguMgSvvRbJ2miMl2krUhfUhffPbq+klMJgd1jnX3pc7RnzM7P22IvFpMPkEYmX5JqFMLLed4KMkgSYTwT6SESh74RMr1htW/l3N4tbmyggyWWWVCfFOQQvCjNxeQ6+GE7eT/4FvrNTftO06wM2722vmzVcsK02QNi2HXNmdiKLs+k4lKmbDoyb/u1cmrzUGZ00sbTnK+2Em5nVVWRcgGggsQqjSJLlconwkk7FDAicyezW7MBjUmkSppLPcByXJi0pjWFlXcqlaSTKeQT7l+/PDrVhP48lBNp8myWaziXNTx9blyeDnE8wrfTGNJ3ERSs0YArnSv1/N+2x7sH6DeU4uVjbpTBLmUAUmCC6v+0IgAt4qPfxJPI0nF07G0wnHrawPJZoR0uEawb6VXWPRiKNb+oCS0SX2fxs8hu0U6GtIKo3KcM0d/81wAB9gdwB5TjmsMj0ZJdiWPwxX/yICPm4cpWmW6Omj023Ly19qcqwPiwjW6BQuuSM8zyjPzxmkK9SevsPkL3fDVPGrOB17zta1OzuGL+o16prfIIJH2p6ORd2s8zNg9iPRT4AYeGsXbbSTAmal//hHn+6CNWjkfKLkDL33uJKSqvBd27H2vr25vbm/gQiWw/BV2oy9HG54isNND9V2BHRS2a01oNHkNeZs20dlYcJBQPeSxaCi3QPXdZaMY8nrS+FfLuysS9P0gti2T/b9p2c9nnS2t+beLNq2/RcDwzYg","sidebar_class_name":"delete api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Create an Organization","permalink":"/docs/1backend/create-organization"},"next":{"title":"Add a User to an Organization","permalink":"/docs/1backend/add-user-to-organization"}}');var a=i(74848),o=i(28453),r=i(53746),s=i.n(r),c=i(56518),d=i.n(c),l=i(99972),p=i.n(l),m=i(25342),u=i.n(m),g=(i(44215),i(82223),i(24861));const b={id:"remove-user-from-organization",title:"Remove a User from an Organization",description:"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization.",sidebar_label:"Remove a User from an Organization",hide_title:!0,hide_table_of_contents:!0,api:"eJzNVk1v2zgQ/SvEnHYBRXLqFljotMk2XXg32BZxekp9GFNjiQ1FqiRl1xX03xcjybH8kQLtAsFebImajzdvyMdpICMvnaqCsgZSuNLabrxAI7AOhXXqG2Wi9uREsAKzTKCxoSC3XxO+IqlWSgrrcjTqG3KoWNwX1BttlNZiSQK9V7mhbOzirCaxUaFQRoSCDkNABLYi173MMkjBUWnX9NGTe+ds+X5kCxFU6LCkQM5D+tAclTW2FbO3EIHi5QpDAREYLAlSGCefZRCBoy+1cpRBGlxNEXhZUImQNhC2FXv44JTJoW2j44QM8tlETMsPJlj0xuTDtc22bCGtCWQCP2JVaSU73Mlnz/mb01B2+ZlkgLZtoyOsdx2tooPMxIoDtu76tNAyAl9Z48lz2FeTCf+dKRuzjDLhaynJ+1Wt9Rai/wC3jeD1uVwzs0atMvHX/P0/P5KgcryrgurLIOesO9vSZ5Bcnqna7A/LCyKZniJ5Z91SZRmZF4Tx+hTGeAMl3Z4wNoiVrc3L8fPm/J4J5AxqMSe3JiduupgvA6mNwJOsnQrbTqGuCR25qzoUkD4s+IQHzFm8+lM0X0tYRFBSKCyLX0aaAnVCxx6QsIpc+LVMxrqVNIcq1nZmSdNLTguMgSvvRbJ2miMl2krUhfUhffPbq+klMJgd1jnX3pc7RnzM7P22IvFpMPkEYmX5JqFMLLed4KMkgSYTwT6SESh74RMr1htW/l3N4tbmyggyWWWVCfFOQQvCjNxeQ6+GE7eT/4FvrNTftO06wM2722vmzVcsK02QNi2HXNmdiKLs+k4lKmbDoyb/u1cmrzUGZ00sbTnK+2Em5nVVWRcgGggsQqjSJLlconwkk7FDAicyezW7MBjUmkSppLPcByXJi0pjWFlXcqlaSTKeQT7l+/PDrVhP48lBNp8myWaziXNTx9blyeDnE8wrfTGNJ3ERSs0YArnSv1/N+2x7sH6DeU4uVjbpTBLmUAUmCC6v+0IgAt4qPfxJPI0nF07G0wnHrawPJZoR0uEawb6VXWPRiKNb+oCS0SX2fxs8hu0U6GtIKo3KcM0d/81wAB9gdwB5TjmsMj0ZJdiWPwxX/yICPm4cpWmW6Omj023Ly19qcqwPiwjW6BQuuSM8zyjPzxmkK9SevsPkL3fDVPGrOB17zta1OzuGL+o16prfIIJH2p6ORd2s8zNg9iPRT4AYeGsXbbSTAmal//hHn+6CNWjkfKLkDL33uJKSqvBd27H2vr25vbm/gQiWw/BV2oy9HG54isNND9V2BHRS2a01oNHkNeZs20dlYcJBQPeSxaCi3QPXdZaMY8nrS+FfLuysS9P0gti2T/b9p2c9nnS2t+beLNq2/RcDwzYg",sidebar_class_name:"delete api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},h=void 0,z={},S=[];function f(e){const n={p:"p",...(0,o.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(g.default,{as:"h1",className:"openapi__heading",children:"Remove a User from an Organization"}),"\n",(0,a.jsx)(s(),{method:"delete",path:"/user-svc/organization/{organizationId}/user/{userId}",context:"endpoint"}),"\n",(0,a.jsx)(n.p,{children:"Allows an authorized user to add another user to a specific organization. The user will be assigned a specific role within the organization."}),"\n",(0,a.jsx)(g.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,a.jsx)(d(),{parameters:[{description:"Organization ID",in:"path",name:"organizationId",required:!0,schema:{type:"string"}},{description:"User ID",in:"path",name:"userId",required:!0,schema:{type:"string"}}]}),"\n",(0,a.jsx)(p(),{title:"Body",body:{content:{"application/json":{schema:{type:"object"}}},description:"Remove User From Organization Request"}}),"\n",(0,a.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"User added successfully",content:{"application/json":{schema:{type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},403:{description:"Forbidden",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},404:{description:"Organization/User not found",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function j(e={}){const{wrapper:n}={...(0,o.R)(),...e.components};return n?(0,a.jsx)(n,{...e,children:(0,a.jsx)(f,{...e})}):f(e)}}}]);