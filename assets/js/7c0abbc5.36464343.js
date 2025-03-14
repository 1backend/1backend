"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[7644],{16335:(e,t,i)=>{i.r(t),i.d(t,{assets:()=>k,contentTitle:()=>y,default:()=>C,frontMatter:()=>m,metadata:()=>o,toc:()=>v});const o=JSON.parse('{"id":"openorch/get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","source":"@site/docs/openorch/get-public-key.api.mdx","sourceDirName":"openorch","slug":"/openorch/get-public-key","permalink":"/docs/openorch/get-public-key","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"get-public-key","title":"Get Public Key","description":"Get the public key to parse and verify the JWT.","sidebar_label":"Get Public Key","hide_title":true,"hide_table_of_contents":true,"api":"eJy9VE2P0zAQ/SvWnNOk7IIEOVEktFpAdKXuisNuD1Nnmph1bDN2WkqU/44madkPuMCBU6zJ+PnNmzfTQ0VRswnJeAclXFBSqSEVuo01Wt3TQSWvAnIkha5SO2KzPYwpH75c55CBD8Qoty8rKKGmdDVe/UgHyIApBu8iRSh7OJvP5fP0weVHyEB7l8gl+YshWKNHwOJrlJQeom6oRTkFlueSmQDDr5fKHtIhEJQQExtXwzBkp4jffCWdYBgk9vJPHC7dDq2p1IfV8rPyrFoTo3G1CsTj0Ttlqn+nScye/4bii98p3jjsUuPZ/KD/xUSCWEcob+EmEqvVTsM6g5ZS44+thgwCpgZKKLpIPIs7XUxNmd2P/Y/EO2LB6KFjK4mF9Rpt42MqX70+O38Bw1rydMcmHVZCfOL6jpCJF53AP1fj+hBI3R1T7kBtvbV+T5XaHBSqGFBPbk3+npxC/a0zTJXasm9H657qUZ98bZwiVwVvXBI7G8FvCCtiyMBhK6osjtqPKsMvsTAYMZ8oZdzWC09pDOqxMdSikYojWopvxVCdxcTe5dq3j7CvLtWqC8GzyDmJ1KQUyqLwgZxn3eSe6wKG7JkKC2XR1R3WNMPa+ZiMVq3R7EV1oymqLWNLe8/3autZbTpjK/H14lI9ck2Usq3R5CIJ7xOvi6tPaneez5+wimVR7Pf7vHbdyOp4LxZYBzs7z+d5k1orXBNxG5fb1cTloai4x7omzo0vxpRC9DTJSsoykFuybiADsc1U5jw/z+cz1vnZG8ENPqYW3SOmsrOmpaOmrfNEpv5hWP5hux07neh7KoJF44TCKEd/tP4tnKwv0/Bg/nUGYnJJ6PsNRrphOwwS/tYRH6C8XWewQza4kdpv10N28p1Mi0CUsNCagvhih7abLPds2ofHI3nx/hqG4SdpQwBt","sidebar_class_name":"get api-method","info_path":"docs/1backend/1backend","custom_edit_url":null},"sidebar":"tutorialSidebar","previous":{"title":"Get Permissions by Role","permalink":"/docs/openorch/get-permissions-by-role"},"next":{"title":"Get all Roles","permalink":"/docs/openorch/get-roles"}}');var s=i(74848),n=i(28453),p=i(53746),r=i.n(p),a=i(56518),c=i.n(a),d=i(99972),l=i.n(d),h=i(25342),u=i.n(h),b=(i(44215),i(82223),i(24861));const m={id:"get-public-key",title:"Get Public Key",description:"Get the public key to parse and verify the JWT.",sidebar_label:"Get Public Key",hide_title:!0,hide_table_of_contents:!0,api:"eJy9VE2P0zAQ/SvWnNOk7IIEOVEktFpAdKXuisNuD1Nnmph1bDN2WkqU/44madkPuMCBU6zJ+PnNmzfTQ0VRswnJeAclXFBSqSEVuo01Wt3TQSWvAnIkha5SO2KzPYwpH75c55CBD8Qoty8rKKGmdDVe/UgHyIApBu8iRSh7OJvP5fP0weVHyEB7l8gl+YshWKNHwOJrlJQeom6oRTkFlueSmQDDr5fKHtIhEJQQExtXwzBkp4jffCWdYBgk9vJPHC7dDq2p1IfV8rPyrFoTo3G1CsTj0Ttlqn+nScye/4bii98p3jjsUuPZ/KD/xUSCWEcob+EmEqvVTsM6g5ZS44+thgwCpgZKKLpIPIs7XUxNmd2P/Y/EO2LB6KFjK4mF9Rpt42MqX70+O38Bw1rydMcmHVZCfOL6jpCJF53AP1fj+hBI3R1T7kBtvbV+T5XaHBSqGFBPbk3+npxC/a0zTJXasm9H657qUZ98bZwiVwVvXBI7G8FvCCtiyMBhK6osjtqPKsMvsTAYMZ8oZdzWC09pDOqxMdSikYojWopvxVCdxcTe5dq3j7CvLtWqC8GzyDmJ1KQUyqLwgZxn3eSe6wKG7JkKC2XR1R3WNMPa+ZiMVq3R7EV1oymqLWNLe8/3autZbTpjK/H14lI9ck2Usq3R5CIJ7xOvi6tPaneez5+wimVR7Pf7vHbdyOp4LxZYBzs7z+d5k1orXBNxG5fb1cTloai4x7omzo0vxpRC9DTJSsoykFuybiADsc1U5jw/z+cz1vnZG8ENPqYW3SOmsrOmpaOmrfNEpv5hWP5hux07neh7KoJF44TCKEd/tP4tnKwv0/Bg/nUGYnJJ6PsNRrphOwwS/tYRH6C8XWewQza4kdpv10N28p1Mi0CUsNCagvhih7abLPds2ofHI3nx/hqG4SdpQwBt",sidebar_class_name:"get api-method",info_path:"docs/1backend/1backend",custom_edit_url:null},y=void 0,k={},v=[];function P(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(b.default,{as:"h1",className:"openapi__heading",children:"Get Public Key"}),"\n",(0,s.jsx)(r(),{method:"get",path:"/user-svc/public-key",context:"endpoint"}),"\n",(0,s.jsx)(t.p,{children:"Get the public key to parse and verify the JWT."}),"\n",(0,s.jsx)(c(),{parameters:void 0}),"\n",(0,s.jsx)(l(),{title:"Body",body:void 0}),"\n",(0,s.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"OK",content:{"application/json":{schema:{properties:{publicKey:{type:"string"}},type:"object"}}}},400:{description:"Invalid JSON or missing permission id",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function C(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,s.jsx)(t,{...e,children:(0,s.jsx)(P,{...e})}):P(e)}}}]);