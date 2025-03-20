"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7332],{77203:(e,s,i)=>{i.r(s),i.d(s,{assets:()=>g,contentTitle:()=>b,default:()=>w,frontMatter:()=>P,metadata:()=>t,toc:()=>y});const t=JSON.parse('{"id":"1backend/get-permissions-by-role","title":"Get Permissions by Role","description":"Retrieve permissions associated with a specific role ID.","source":"@site/docs/1backend/get-permissions-by-role.api.mdx","sourceDirName":"1backend","slug":"/1backend/get-permissions-by-role","permalink":"/docs/1backend/get-permissions-by-role","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-permissions-by-role","title":"Get Permissions by Role","description":"Retrieve permissions associated with a specific role ID.","sidebar_label":"Get Permissions by Role","hide_title":true,"hide_table_of_contents":true,"api":"eJylVt+P2zYM/lcIvawFfPZdbwW2PO36A0XWdi2aXoHhnAdGZmz1ZMmV6GRZ4P99oJVccndpgaEviW2R/D5S5CdtVUVRB9Ox8U5N1CfiYGhF0FFoTYzGuwgYo9cGmSpYG24AIXakzdJoCN4STF/lKlO+o4ASZlqpiaqJPx5CvNh88pZUpjoM2BJTiGpys30InoKpTBl57ZAblSmHLamJEqRppTIV6FtvAlVqwqGnTEXdUItqslW86cQycjCuVsMwF+PYeRcpyvqz83P5uw/64a3KlPaOybGsYtdZo8dEiq9RTLZHEF2QNNmkgEdFklfD1MbHVjqQ1O6KT3DM7rM5sW6qx5ypzqFUfaSQrwytKZRKZY9dU+VOO19HCvDlB85+7ShMT4DPKKyMJlg3HvzaReDmuF1KV7prpykwGgdmCdyYCGhN7WJqH7GPm8jU/hKhjwR+CdH2dcxLN42xJyCnfe+YAlUTmEIgyaSCV17fUoDZSgN7eOmdQOw+GAe89hCZughPFhtoTWS8pad56T43BBKhEBRosAK0gbDagG7Q1VRJuHJsgjHgWVzprFSw6BkqZIQ1RohsrIUFGVdDxBVVpRPMhqBU1UhsttIvA1Xk2KCNpQLGhaW8dFdLppCyEG/xGZdksmBN1mYwhcYwoLBA5zxD31XIBL2TfaiO66uAQvChdAvSOJZPdio2poNA1lAE71LRl4ZsBQG5oQDcYKIrbTPWO4e/37//kp/a/IR+umWHO3u/+Eqa1eEDhoCbUxaDfPv11PhN3QqtqeDP2Ye//s8gPiSVAC4eA1w77LnxwfxL1c8CPD+dAVNwaEEGgwK8lt35OaQhU5F0HwxvRpl8QRgoXPXcqMnNXHSNsRYFTWM8W2k1z1RL3Pid9o5SK+aqkP2Wji5EQYtt0tGhONYugRPuSZT7YMWvsF6jbXzkyfPfnl1eKMHd05oJ+yRvx+Qe1ubzppP5SCalgqW31q+pgsVmPEVQE6CT+bslB6iTssMy+Hbs1X168M7XxgG5qvPGcb4/IhrCisLhkLjabfVY6kNjY2feUmpM45Z+FGUZdj3uDrVoJOOIluIf0bi6t8jBu1z79ij2xynM+q7zQaqbitQwd5OiuFigviVXiUPxSNTV1fTMIZsVQWt08DEJaITOIi99aCUdazS5OMr1Hu/Nx3ewuszP76HFSVGs1+u8dn3uQ13s/GKBdWfPLvPzvOHWjjNJoY0flju5PpCNa6xrCrnxxWhSSJ0MWzG5eJESUZmSdkj0z/PL/Pws6PzZ7xK385FbdEdM3xDD0WEvm7s77h8cbncj8TP3jN2eMv3DRWfRuFGwpEDbXc/fqH3Py3UhMZnc3R+OG3+eKWlwcdluFxjpOthhkM/fegoyfPNMrTAYket0YzFRnis1WaKN9IMUn3zaXVSewuFic5L8vkvdRuqOtpc3lalb2hwuPsN8yPb9LkTS4pXW1PGR2yOpGY6V4c3rzypTuJvVw3RIsGz/INFPUno4XYmC/A7Zd1y22zR7w3Bnn5a+63E30slaKjQfhuE/prmuhA==","sidebar_class_name":"get api-method","info_path":"docs/1backend/1-backend","custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"Delete a Role","permalink":"/docs/1backend/delete-role"},"next":{"title":"Set Role Permissions","permalink":"/docs/1backend/set-role-permission"}}');var n=i(74848),a=i(28453),o=i(53746),r=i.n(o),d=i(56518),p=i.n(d),l=i(99972),c=i.n(l),h=i(25342),u=i.n(h),m=(i(44215),i(82223),i(24861));const P={id:"get-permissions-by-role",title:"Get Permissions by Role",description:"Retrieve permissions associated with a specific role ID.",sidebar_label:"Get Permissions by Role",hide_title:!0,hide_table_of_contents:!0,api:"eJylVt+P2zYM/lcIvawFfPZdbwW2PO36A0XWdi2aXoHhnAdGZmz1ZMmV6GRZ4P99oJVccndpgaEviW2R/D5S5CdtVUVRB9Ox8U5N1CfiYGhF0FFoTYzGuwgYo9cGmSpYG24AIXakzdJoCN4STF/lKlO+o4ASZlqpiaqJPx5CvNh88pZUpjoM2BJTiGpys30InoKpTBl57ZAblSmHLamJEqRppTIV6FtvAlVqwqGnTEXdUItqslW86cQycjCuVsMwF+PYeRcpyvqz83P5uw/64a3KlPaOybGsYtdZo8dEiq9RTLZHEF2QNNmkgEdFklfD1MbHVjqQ1O6KT3DM7rM5sW6qx5ypzqFUfaSQrwytKZRKZY9dU+VOO19HCvDlB85+7ShMT4DPKKyMJlg3HvzaReDmuF1KV7prpykwGgdmCdyYCGhN7WJqH7GPm8jU/hKhjwR+CdH2dcxLN42xJyCnfe+YAlUTmEIgyaSCV17fUoDZSgN7eOmdQOw+GAe89hCZughPFhtoTWS8pad56T43BBKhEBRosAK0gbDagG7Q1VRJuHJsgjHgWVzprFSw6BkqZIQ1RohsrIUFGVdDxBVVpRPMhqBU1UhsttIvA1Xk2KCNpQLGhaW8dFdLppCyEG/xGZdksmBN1mYwhcYwoLBA5zxD31XIBL2TfaiO66uAQvChdAvSOJZPdio2poNA1lAE71LRl4ZsBQG5oQDcYKIrbTPWO4e/37//kp/a/IR+umWHO3u/+Eqa1eEDhoCbUxaDfPv11PhN3QqtqeDP2Ye//s8gPiSVAC4eA1w77LnxwfxL1c8CPD+dAVNwaEEGgwK8lt35OaQhU5F0HwxvRpl8QRgoXPXcqMnNXHSNsRYFTWM8W2k1z1RL3Pid9o5SK+aqkP2Wji5EQYtt0tGhONYugRPuSZT7YMWvsF6jbXzkyfPfnl1eKMHd05oJ+yRvx+Qe1ubzppP5SCalgqW31q+pgsVmPEVQE6CT+bslB6iTssMy+Hbs1X168M7XxgG5qvPGcb4/IhrCisLhkLjabfVY6kNjY2feUmpM45Z+FGUZdj3uDrVoJOOIluIf0bi6t8jBu1z79ij2xynM+q7zQaqbitQwd5OiuFigviVXiUPxSNTV1fTMIZsVQWt08DEJaITOIi99aCUdazS5OMr1Hu/Nx3ewuszP76HFSVGs1+u8dn3uQ13s/GKBdWfPLvPzvOHWjjNJoY0flju5PpCNa6xrCrnxxWhSSJ0MWzG5eJESUZmSdkj0z/PL/Pws6PzZ7xK385FbdEdM3xDD0WEvm7s77h8cbncj8TP3jN2eMv3DRWfRuFGwpEDbXc/fqH3Py3UhMZnc3R+OG3+eKWlwcdluFxjpOthhkM/fegoyfPNMrTAYket0YzFRnis1WaKN9IMUn3zaXVSewuFic5L8vkvdRuqOtpc3lalb2hwuPsN8yPb9LkTS4pXW1PGR2yOpGY6V4c3rzypTuJvVw3RIsGz/INFPUno4XYmC/A7Zd1y22zR7w3Bnn5a+63E30slaKjQfhuE/prmuhA==",sidebar_class_name:"get api-method",info_path:"docs/1backend/1-backend",custom_edit_url:null},b=void 0,g={},y=[];function R(e){const s={p:"p",...(0,a.R)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Get Permissions by Role"}),"\n",(0,n.jsx)(r(),{method:"get",path:"/user-svc/role/{roleId}/permissions",context:"endpoint"}),"\n",(0,n.jsx)(s.p,{children:"Retrieve permissions associated with a specific role ID."}),"\n",(0,n.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,n.jsx)(p(),{parameters:[{description:"Role ID",in:"path",name:"roleId",required:!0,schema:{type:"string"}}]}),"\n",(0,n.jsx)(c(),{title:"Body",body:void 0}),"\n",(0,n.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{permissions:{items:{properties:{createdAt:{type:"string"},description:{type:"string"},id:{description:'eg. "user.viewer"',type:"string"},name:{description:'eg. "User Viewer"',type:"string"},ownerId:{description:'Service who owns the permission\n\nUncertain if this aligns with the system\'s use of slugs.\nIssue encountered: I renamed Docker Svc to Container Svc in two steps (by mistake).\nThe name/slug had already changed to "container-svc," but data was still being saved\nin the "dockerSvcCredentials" table.\nAfter renaming the tables as well, I hit a "cannot update unowned permission" error\nbecause ownership relies on this field rather than the user slug. YMMV.',type:"string"},updatedAt:{type:"string"}},type:"object"},type:"array"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{type:"string"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{type:"string"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{type:"string"}}}}}})]})}function w(e={}){const{wrapper:s}={...(0,a.R)(),...e.components};return s?(0,n.jsx)(s,{...e,children:(0,n.jsx)(R,{...e})}):R(e)}}}]);