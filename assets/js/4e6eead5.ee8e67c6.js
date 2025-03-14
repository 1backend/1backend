"use strict";(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[1770],{44587:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>k,contentTitle:()=>C,default:()=>g,frontMatter:()=>f,metadata:()=>r,toc:()=>b});const r=JSON.parse('{"id":"openorch/checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","source":"@site/docs/openorch/checkout-repo.api.mdx","sourceDirName":"openorch","slug":"/openorch/checkout-repo","permalink":"/docs/openorch/checkout-repo","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","frontMatter":{"id":"checkout-repo","title":"Checkout a git repository","description":"Checkout a git repository over https or ssh at a specific version into a temporary directory.","sidebar_label":"Checkout a git repository","hide_title":true,"hide_table_of_contents":true,"api":"eJzVVt9v2zYQ/lcIPnWAIqXNCmx6WtJ1a7ZiCeL0KTEKmjpLbCiSPVJWNUP/+3CUZCuxUxR9KLAXQz4e77779R23vAAvUbmgrOE5f1OBfLBNYIKVKjAEZ70KFjtmN4CsCsF5ZpF5XzFBWt6BVGsl2QbQK2uYMsEywQLUzqLAjhUKQZKJ9N5cA64t1p4uVkJr2zKprQHWqlCxWhlVC80q5aPLtUW2Fj4AMjniSnnCrQMUhPey4DmfTm7AWZ5whM8N+HBhi47nWy6tCWACfQrntJLxYvbJU7Rb7mUFtaAvh2Q2KPDxn/C+tVjQ9+MEXY8nlIRgH8BElO9ub68XTDSh4gkPnQOecx9QmZL3Cfe++vgA3aGxxeIdc6g2IgB7gI69sPFE6GiUTqU1BiQJf/qK5Y+u/RrUydjclVozMBI7F6DY+z3qJIZ5aP72SPSHERy12KA+tPdHo/W84T7cvGcvIC3TZGi7PMtKFapmlUpbZ40HzEj7uAMPaEQNh14+jCcz4GN26M434h97/dD6BQojq4QFUSZkV9q6VoEt3p0fmul3Erv6BDLwnkTPjCN1N7sZenvscoVQ8DxgAz0JvLPGD9376vT0SK81UoL360brbhgnKBiZDhXM0s6T752ZQuGh19+n4WdtBQhPnLFW+DmWb0tSn/Cfj0V4aTZCq4L9tbj65/vDAEQbA/lWJC+PNJmhYbCo/oXihyF5fTwngfpdswUgEfjbaPPHQCJ+AtmgCh3P77b8AgQCnhNJ5nfLfpnwIErP8zu+sA1KYIuN5MuE1xAqS9TubGx2J+gGz3zUOvEbGQc/m5ifkx+Kzkc3kVx4lmkrha6sD/nrX16dveTkcMKzoPiGkOaoDiiuc8DuR5V7ztaWVhYUbNXF1SckMGGKcREIOQwlW6OtY6cT2VBU7L0tlWFgCmeViUtMkf0KRAFUjoGr+PnYNbEY+2EQTv0NXcwyFehmv+LefhG10/B4ZU1VmS2eA9GwMfbikeP3giGL+787Qt3Ldiw4Y0Zl1nbau0LG/oJaqGhLaPC/eWXKRouA1hCPz2K/vmSLxjmLYeeeeD/PMuvAWJRVarHM+AFJnjMtTNmIEk5EaawPSrJaSbTUFUqCZ2sUNbQWHyKrrxqlC2VKdn7JZs3vqSxaSTA+JnTC9ef1e7Y5S08foaJt1LZtWpomohrv+UyUTp+cpadpFWodtydg7a/WiwHLPijfirIETJXNokpGVVCBismvHJgrlNWjJJ+mZ+npCcr01a9kl4ajFmaG9Nm3G3+Ssdmz6H/x4BvHIMCXkDktlJm9IQZyuON7cojrMb4FdwSxTDgRAelttyvh4QPqvifx5waQ2GmZ8I1AJVaU/7tln0yzSYwyjNCbIWcnRApUGKGbYTif0GefTDfOpQQXvqo7p7vrq8UtT/hqfLrWtqA7KFqKSLQ85/HxG3uVFKJsy6fm5zkfbBJPiJHP9gxCkJLpg6Kajkw3Q/iUgYZA6JfCOnplux34qe93+sPRszd2tDdoUz2Xfd//B6DkXyc=","sidebar_class_name":"post api-method","info_path":"docs/1backend/1backend","custom_edit_url":null},"sidebar":"tutorialSidebar","previous":{"title":"Check","permalink":"/docs/openorch/check"},"next":{"title":"Get Container Daemon Information","permalink":"/docs/openorch/container-daemon-info"}}');var i=o(74848),n=o(28453),s=o(53746),a=o.n(s),c=o(56518),p=o.n(c),h=o(99972),d=o.n(h),l=o(25342),u=o.n(l),m=(o(44215),o(82223),o(24861));const f={id:"checkout-repo",title:"Checkout a git repository",description:"Checkout a git repository over https or ssh at a specific version into a temporary directory.",sidebar_label:"Checkout a git repository",hide_title:!0,hide_table_of_contents:!0,api:"eJzVVt9v2zYQ/lcIPnWAIqXNCmx6WtJ1a7ZiCeL0KTEKmjpLbCiSPVJWNUP/+3CUZCuxUxR9KLAXQz4e77779R23vAAvUbmgrOE5f1OBfLBNYIKVKjAEZ70KFjtmN4CsCsF5ZpF5XzFBWt6BVGsl2QbQK2uYMsEywQLUzqLAjhUKQZKJ9N5cA64t1p4uVkJr2zKprQHWqlCxWhlVC80q5aPLtUW2Fj4AMjniSnnCrQMUhPey4DmfTm7AWZ5whM8N+HBhi47nWy6tCWACfQrntJLxYvbJU7Rb7mUFtaAvh2Q2KPDxn/C+tVjQ9+MEXY8nlIRgH8BElO9ub68XTDSh4gkPnQOecx9QmZL3Cfe++vgA3aGxxeIdc6g2IgB7gI69sPFE6GiUTqU1BiQJf/qK5Y+u/RrUydjclVozMBI7F6DY+z3qJIZ5aP72SPSHERy12KA+tPdHo/W84T7cvGcvIC3TZGi7PMtKFapmlUpbZ40HzEj7uAMPaEQNh14+jCcz4GN26M434h97/dD6BQojq4QFUSZkV9q6VoEt3p0fmul3Erv6BDLwnkTPjCN1N7sZenvscoVQ8DxgAz0JvLPGD9376vT0SK81UoL360brbhgnKBiZDhXM0s6T752ZQuGh19+n4WdtBQhPnLFW+DmWb0tSn/Cfj0V4aTZCq4L9tbj65/vDAEQbA/lWJC+PNJmhYbCo/oXihyF5fTwngfpdswUgEfjbaPPHQCJ+AtmgCh3P77b8AgQCnhNJ5nfLfpnwIErP8zu+sA1KYIuN5MuE1xAqS9TubGx2J+gGz3zUOvEbGQc/m5ifkx+Kzkc3kVx4lmkrha6sD/nrX16dveTkcMKzoPiGkOaoDiiuc8DuR5V7ztaWVhYUbNXF1SckMGGKcREIOQwlW6OtY6cT2VBU7L0tlWFgCmeViUtMkf0KRAFUjoGr+PnYNbEY+2EQTv0NXcwyFehmv+LefhG10/B4ZU1VmS2eA9GwMfbikeP3giGL+787Qt3Ldiw4Y0Zl1nbau0LG/oJaqGhLaPC/eWXKRouA1hCPz2K/vmSLxjmLYeeeeD/PMuvAWJRVarHM+AFJnjMtTNmIEk5EaawPSrJaSbTUFUqCZ2sUNbQWHyKrrxqlC2VKdn7JZs3vqSxaSTA+JnTC9ef1e7Y5S08foaJt1LZtWpomohrv+UyUTp+cpadpFWodtydg7a/WiwHLPijfirIETJXNokpGVVCBismvHJgrlNWjJJ+mZ+npCcr01a9kl4ajFmaG9Nm3G3+Ssdmz6H/x4BvHIMCXkDktlJm9IQZyuON7cojrMb4FdwSxTDgRAelttyvh4QPqvifx5waQ2GmZ8I1AJVaU/7tln0yzSYwyjNCbIWcnRApUGKGbYTif0GefTDfOpQQXvqo7p7vrq8UtT/hqfLrWtqA7KFqKSLQ85/HxG3uVFKJsy6fm5zkfbBJPiJHP9gxCkJLpg6Kajkw3Q/iUgYZA6JfCOnplux34qe93+sPRszd2tDdoUz2Xfd//B6DkXyc=",sidebar_class_name:"post api-method",info_path:"docs/1backend/1backend",custom_edit_url:null},C=void 0,k={},b=[];function y(e){const t={p:"p",...(0,n.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(m.default,{as:"h1",className:"openapi__heading",children:"Checkout a git repository"}),"\n",(0,i.jsx)(a(),{method:"post",path:"/source-svc/repo/checkout",context:"endpoint"}),"\n",(0,i.jsx)(t.p,{children:"Checkout a git repository over https or ssh at a specific version into a temporary directory.\nPerforms a shallow clone with minimal history for faster checkout."}),"\n",(0,i.jsx)(m.default,{id:"request",as:"h2",className:"openapi-tabs__heading",children:"Request"}),"\n",(0,i.jsx)(p(),{parameters:void 0}),"\n",(0,i.jsx)(d(),{title:"Body",body:{content:{"application/json":{schema:{properties:{password:{description:"Password or token for HTTPS auth",type:"string"},ssh_key:{description:"SSH private key (optional for SSH connection)",type:"string"},ssh_key_pwd:{description:"Password for SSH private key if encrypted (optional)",type:"string"},token:{description:"Token for HTTPS auth (optional for SSH)",type:"string"},url:{description:"Full repository URL (e.g., https://github.com/user/repo)",type:"string"},username:{description:"Username for HTTPS or SSH user (optional for SSH)",type:"string"},version:{description:"Branch, tag, or commit SHA",type:"string"}},type:"object"}}},description:"Checkout Repo Request",required:!0}}),"\n",(0,i.jsx)(u(),{id:void 0,label:void 0,responses:{200:{description:"Successfully checked out the repository",content:{"application/json":{schema:{properties:{dir:{description:"Directory where the repository was checked out",type:"string"}},type:"object"}}}},400:{description:"Invalid JSON",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},401:{description:"Unauthorized",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}},500:{description:"Internal Server Error",content:{"application/json":{schema:{properties:{error:{type:"string"}},type:"object"}}}}}})]})}function g(e={}){const{wrapper:t}={...(0,n.R)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(y,{...e})}):y(e)}}}]);