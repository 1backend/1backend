(self.webpackChunksingulatron_api_docs=self.webpackChunksingulatron_api_docs||[]).push([[9557],{95196:function(e,t,n){"use strict";var a=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0});const o=a(n(96540)),r=a(n(4213));t.default=function(e){let{url:t,proxy:n}=e;return o.default.createElement("div",{style:{float:"right"},className:"dropdown dropdown--hoverable dropdown--right"},o.default.createElement("button",{className:"export-button button button--sm button--secondary"},"Export"),o.default.createElement("ul",{className:"export-dropdown dropdown__menu"},o.default.createElement("li",null,o.default.createElement("a",{onClick:e=>{e.preventDefault(),(e=>{let t;(e.endsWith("json")||e.endsWith("yaml")||e.endsWith("yml"))&&(t=e.substring(e.lastIndexOf("/")+1)),r.default.saveAs(e,t||"openapi.txt")})(`${t}`)},className:"dropdown__link",href:`${t}`},"OpenAPI Spec"))))}},98205:function(e,t,n){"use strict";var a=this&&this.__importDefault||function(e){return e&&e.__esModule?e:{default:e}};Object.defineProperty(t,"__esModule",{value:!0}),t.default=function(e){const{colorMode:t}=(0,r.useColorMode)(),{logo:n,darkLogo:a}=e,c=()=>"dark"===t?a?.altText??n?.altText:n?.altText,l=(0,i.default)(n?.url),d=(0,i.default)(a?.url);if(n&&a)return o.default.createElement(s.default,{alt:c(),sources:{light:l,dark:d},className:"openapi__logo"});if(n||a)return o.default.createElement(s.default,{alt:c(),sources:{light:l??d,dark:l??d},className:"openapi__logo"});return};const o=a(n(96540)),r=n(87113),i=a(n(76813)),s=a(n(48110))},18247:(e,t,n)=>{"use strict";n.r(t),n.d(t,{assets:()=>f,contentTitle:()=>m,default:()=>g,frontMatter:()=>p,metadata:()=>a,toc:()=>h});const a=JSON.parse('{"id":"1backend/1-backend","title":"1Backend","description":"A common backend for your AI applications\u2014microservices-based and built to scale.","source":"@site/docs/1backend/1-backend.info.mdx","sourceDirName":"1backend","slug":"/1backend/1-backend","permalink":"/docs/1backend/1-backend","draft":false,"unlisted":false,"editUrl":null,"tags":[],"version":"current","sidebarPosition":0,"frontMatter":{"id":"1-backend","title":"1Backend","description":"A common backend for your AI applications\u2014microservices-based and built to scale.","sidebar_label":"Introduction","sidebar_position":0,"hide_title":true,"custom_edit_url":null},"sidebar":"openApiSidebar","previous":{"title":"1Backend","permalink":"/docs/category/1backend-api"},"next":{"title":"Events","permalink":"/docs/1backend/events"}}');var o=n(74848),r=n(28453),i=(n(98205),n(24861)),s=n(79569),c=n.n(s),l=n(82223),d=n(95196),u=n.n(d);const p={id:"1-backend",title:"1Backend",description:"A common backend for your AI applications\u2014microservices-based and built to scale.",sidebar_label:"Introduction",sidebar_position:0,hide_title:!0,custom_edit_url:null},m=void 0,f={},h=[];function b(e){const t={a:"a",p:"p",...(0,r.R)(),...e.components};return(0,o.jsxs)(o.Fragment,{children:[(0,o.jsx)("span",{className:"theme-doc-version-badge badge badge--secondary",children:"Version: 0.3.0-rc.29"}),"\n",(0,o.jsx)(u(),{url:"https://raw.githubusercontent.com/1backend/1backend/main/server/docs/swagger.yaml",proxy:void 0}),"\n",(0,o.jsx)(i.default,{as:"h1",className:"openapi__heading",children:"1Backend"}),"\n",(0,o.jsx)(t.p,{children:"A common backend for your AI applications\u2014microservices-based and built to scale."}),"\n",(0,o.jsxs)("div",{style:{marginBottom:"2rem"},children:[(0,o.jsx)(i.default,{id:"authentication",as:"h2",className:"openapi-tabs__heading",children:"Authentication"}),(0,o.jsx)(c(),{className:"openapi-tabs__security-schemes",children:(0,o.jsxs)(l.default,{label:"API Key: BearerAuth",value:"BearerAuth",children:[(0,o.jsx)(t.p,{children:'Type "Bearer" followed by a space and token acquired from the User Svc Login endpoint.'}),(0,o.jsx)("div",{children:(0,o.jsx)("table",{children:(0,o.jsxs)("tbody",{children:[(0,o.jsxs)("tr",{children:[(0,o.jsx)("th",{children:(0,o.jsx)(t.p,{children:"Security Scheme Type:"})}),(0,o.jsx)("td",{children:(0,o.jsx)(t.p,{children:"apiKey"})})]}),(0,o.jsxs)("tr",{children:[(0,o.jsx)("th",{children:(0,o.jsx)(t.p,{children:"Header parameter name:"})}),(0,o.jsx)("td",{children:(0,o.jsx)(t.p,{children:"Authorization"})})]})]})})})]})})]}),"\n",(0,o.jsxs)("div",{style:{display:"flex",flexDirection:"column",marginBottom:"var(--ifm-paragraph-margin-bottom)"},children:[(0,o.jsx)("h3",{style:{marginBottom:"0.25rem"},children:(0,o.jsx)(t.p,{children:"Contact"})}),(0,o.jsx)("span",{children:(0,o.jsxs)(t.p,{children:["API Support: ",(0,o.jsx)(t.a,{href:"mailto:sales@singulatron.com",children:"sales@singulatron.com"})]})}),(0,o.jsx)("span",{children:(0,o.jsxs)(t.p,{children:["URL: ",(0,o.jsx)(t.a,{href:"http://1backend.com/",children:"http://1backend.com/"})]})})]}),"\n",(0,o.jsxs)("div",{style:{marginBottom:"var(--ifm-paragraph-margin-bottom)"},children:[(0,o.jsx)("h3",{style:{marginBottom:"0.25rem"},children:(0,o.jsx)(t.p,{children:"Terms of Service"})}),(0,o.jsx)("a",{href:"http://swagger.io/terms/",children:"http://swagger.io/terms/"})]}),"\n",(0,o.jsxs)("div",{style:{marginBottom:"var(--ifm-paragraph-margin-bottom)"},children:[(0,o.jsx)("h3",{style:{marginBottom:"0.25rem"},children:(0,o.jsx)(t.p,{children:"License"})}),(0,o.jsx)("a",{href:"https://www.gnu.org/licenses/agpl-3.0.html",children:(0,o.jsx)(t.p,{children:"AGPL v3.0"})})]})]})}function g(e={}){const{wrapper:t}={...(0,r.R)(),...e.components};return t?(0,o.jsx)(t,{...e,children:(0,o.jsx)(b,{...e})}):b(e)}},4213:function(e,t,n){var a,o,r,i=n(96763);o=[],void 0===(r="function"==typeof(a=function(){"use strict";function t(e,t){return void 0===t?t={autoBom:!1}:"object"!=typeof t&&(i.warn("Deprecated: Expected third argument to be a object"),t={autoBom:!t}),t.autoBom&&/^\s*(?:text\/\S*|application\/xml|\S*\/\S*\+xml)\s*;.*charset\s*=\s*utf-8/i.test(e.type)?new Blob(["\ufeff",e],{type:e.type}):e}function a(e,t,n){var a=new XMLHttpRequest;a.open("GET",e),a.responseType="blob",a.onload=function(){l(a.response,t,n)},a.onerror=function(){i.error("could not download file")},a.send()}function o(e){var t=new XMLHttpRequest;t.open("HEAD",e,!1);try{t.send()}catch(e){}return 200<=t.status&&299>=t.status}function r(e){try{e.dispatchEvent(new MouseEvent("click"))}catch(a){var t=document.createEvent("MouseEvents");t.initMouseEvent("click",!0,!0,window,0,0,0,80,20,!1,!1,!1,!1,0,null),e.dispatchEvent(t)}}var s="object"==typeof window&&window.window===window?window:"object"==typeof self&&self.self===self?self:"object"==typeof n.g&&n.g.global===n.g?n.g:void 0,c=s.navigator&&/Macintosh/.test(navigator.userAgent)&&/AppleWebKit/.test(navigator.userAgent)&&!/Safari/.test(navigator.userAgent),l=s.saveAs||("object"!=typeof window||window!==s?function(){}:"download"in HTMLAnchorElement.prototype&&!c?function(e,t,n){var i=s.URL||s.webkitURL,c=document.createElement("a");t=t||e.name||"download",c.download=t,c.rel="noopener","string"==typeof e?(c.href=e,c.origin===location.origin?r(c):o(c.href)?a(e,t,n):r(c,c.target="_blank")):(c.href=i.createObjectURL(e),setTimeout((function(){i.revokeObjectURL(c.href)}),4e4),setTimeout((function(){r(c)}),0))}:"msSaveOrOpenBlob"in navigator?function(e,n,i){if(n=n||e.name||"download","string"!=typeof e)navigator.msSaveOrOpenBlob(t(e,i),n);else if(o(e))a(e,n,i);else{var s=document.createElement("a");s.href=e,s.target="_blank",setTimeout((function(){r(s)}))}}:function(e,t,n,o){if((o=o||open("","_blank"))&&(o.document.title=o.document.body.innerText="downloading..."),"string"==typeof e)return a(e,t,n);var r="application/octet-stream"===e.type,i=/constructor/i.test(s.HTMLElement)||s.safari,l=/CriOS\/[\d]+/.test(navigator.userAgent);if((l||r&&i||c)&&"undefined"!=typeof FileReader){var d=new FileReader;d.onloadend=function(){var e=d.result;e=l?e:e.replace(/^data:[^;]*;/,"data:attachment/file;"),o?o.location.href=e:location=e,o=null},d.readAsDataURL(e)}else{var u=s.URL||s.webkitURL,p=u.createObjectURL(e);o?o.location=p:location.href=p,o=null,setTimeout((function(){u.revokeObjectURL(p)}),4e4)}});s.saveAs=l.saveAs=l,e.exports=l})?a.apply(t,o):a)||(e.exports=r)}}]);