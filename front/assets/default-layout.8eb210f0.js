import{c as S,R as J,D as K,d as Q,e as X,f as Y,g as Z,h as ee,_ as $,i as k,l as te,r as ae,a as ne,j as oe}from"./index.270778e3.js";/* empty css              */import{d as I,c as i,B as p,aD as v,aG as r,aH as o,u as f,F as d,D as A,aM as se,b8 as ue,a1 as ce,b9 as ie,ba as le,bb as z,bc as N,e as R,w as U,bd as re,C as B,aJ as de,aI as _e,be as pe,o as H,bf as fe,b as me,bg as M,bh as he,bi as ve,a$ as be,aE as F,E as ge,bj as ye,bk as xe,p as De}from"./arco.ea015dc6.js";import{f as V,h as j,i as Ee}from"./vue.198ed283.js";import"./chart.1c717195.js";const L=t=>(z("data-v-17257cfa"),t=t(),N(),t),Le={class:"tag-link"},Ce=L(()=>d("span",null,"\u91CD\u65B0\u52A0\u8F7D",-1)),Ie=L(()=>d("span",null,"\u5173\u95ED\u5F53\u524D\u6807\u7B7E\u9875",-1)),Te=L(()=>d("span",null,"\u5173\u95ED\u5DE6\u4FA7\u6807\u7B7E\u9875",-1)),we=L(()=>d("span",null,"\u5173\u95ED\u53F3\u4FA7\u6807\u7B7E\u9875",-1)),Re=L(()=>d("span",null,"\u5173\u95ED\u5176\u5B83\u6807\u7B7E\u9875",-1)),Be=L(()=>d("span",null,"\u5173\u95ED\u5168\u90E8\u6807\u7B7E\u9875",-1)),Se=I({__name:"tab-item",props:{itemData:{type:Object,default(){return[]}},index:{type:Number,default:0}},setup(t){const a=t,e=V(),n=j(),s=S(),b=u=>{e.push({...u})},c=i(()=>s.getTabList),m=i(()=>a.itemData.fullPath!==n.fullPath),g=i(()=>a.index===0),y=i(()=>[0,1].includes(a.index)),x=i(()=>a.index===c.value.length-1),D=(u,l)=>{if(s.deleteTag(l,u),a.itemData.fullPath===n.fullPath){const h=c.value[l-1];e.push({name:h.name})}},T=()=>c.value.findIndex(u=>u.fullPath===n.fullPath),w=async u=>{const{itemData:l,index:h}=a,E=[...c.value];if(u==="current")D(l,h);else if(u==="left"){const _=T();E.splice(1,a.index-1),s.freshTabList(E),_<h&&e.push({name:l.name})}else if(u==="right"){const _=T();E.splice(a.index+1),s.freshTabList(E),_>h&&e.push({name:l.name})}else if(u==="others"){const _=c.value.filter((P,C)=>C===0||C===a.index);s.freshTabList(_),e.push({name:l.name})}else u==="reload"?(s.deleteCache(l),await e.push({name:J,params:{path:n.fullPath}}),s.addCache(l.name)):(s.resetTabList(),e.push({name:K}))};return(u,l)=>{const h=ce,E=Q,_=ie,P=X,C=Y,O=Z,W=ee,q=le;return p(),v(q,{trigger:"contextMenu","popup-max-height":!1,onSelect:w},{content:r(()=>[o(_,{disabled:f(m),value:"reload"},{default:r(()=>[o(E),Ce]),_:1},8,["disabled","value"]),o(_,{class:"sperate-line",disabled:f(g),value:"current"},{default:r(()=>[o(h),Ie]),_:1},8,["disabled","value"]),o(_,{disabled:f(y),value:"left"},{default:r(()=>[o(P),Te]),_:1},8,["disabled","value"]),o(_,{class:"sperate-line",disabled:f(x),value:"right"},{default:r(()=>[o(C),we]),_:1},8,["disabled","value"]),o(_,{value:"others"},{default:r(()=>[o(O),Re]),_:1},8,["value"]),o(_,{value:"all"},{default:r(()=>[o(W),Be]),_:1},8,["value"])]),default:r(()=>[d("span",{class:A(["arco-tag arco-tag-size-medium arco-tag-checked",{"link-activated":t.itemData.fullPath===u.$route.fullPath}]),onClick:l[1]||(l[1]=G=>b(t.itemData))},[d("span",Le,se(u.$t(t.itemData.title)),1),d("span",{class:"arco-icon-hover arco-tag-icon-hover arco-icon-hover-size-medium arco-tag-close-btn",onClick:l[0]||(l[0]=ue(G=>D(t.itemData,t.index),["stop"]))},[o(h)])],2)]),_:1})}}});const $e=$(Se,[["__scopeId","data-v-17257cfa"]]),ke=t=>(z("data-v-509db106"),t=t(),N(),t),Pe={class:"tab-bar-container"},Me={class:"tab-bar-box"},Fe={class:"tab-bar-scroll"},Ae={class:"tags-wrap"},ze=ke(()=>d("div",{class:"tag-bar-operation"},null,-1)),Ne=I({__name:"index",setup(t){const a=k(),e=S(),n=R(),s=i(()=>e.getTabList),b=i(()=>a.navbar?60:0);return U(()=>a.navbar,()=>{n.value.updatePosition()}),te(c=>{!c.meta.noAffix&&!s.value.some(m=>m.fullPath===c.fullPath)&&e.updateTabList(c)},!0),re(()=>{ae()}),(c,m)=>{const g=pe;return p(),B("div",Pe,[o(g,{ref_key:"affixRef",ref:n,"offset-top":f(b)},{default:r(()=>[d("div",Me,[d("div",Fe,[d("div",Ae,[(p(!0),B(de,null,_e(f(s),(y,x)=>(p(),v($e,{key:y.fullPath,index:x,"item-data":y},null,8,["index","item-data"]))),128))])]),ze])]),_:1},8,["offset-top"])])}}});const Ue=$(Ne,[["__scopeId","data-v-509db106"]]);function He(t,a,e,n=!1){t.addEventListener&&typeof t.addEventListener=="function"&&t.addEventListener(a,e,n)}function Ve(t,a,e,n=!1){t.removeEventListener&&typeof t.removeEventListener=="function"&&t.removeEventListener(a,e,n)}const je=992;function Oe(){return document.body.getBoundingClientRect().width-1<je}function We(t){const a=k();function e(){if(!document.hidden){const s=Oe();a.toggleDevice(s?"mobile":"desktop"),a.toggleMenu(s)}}const n=Ee(e,100);H(()=>{t&&n()}),fe(()=>{He(window,"resize",n)}),me(()=>{Ve(window,"resize",n)})}const qe=I({__name:"page-layout",setup(t){const a=S(),e=i(()=>a.getCacheList);return(n,s)=>{const b=be("router-view");return p(),v(b,null,{default:r(({Component:c,route:m})=>[o(ve,{name:"fade",mode:"out-in",appear:""},{default:r(()=>[m.meta.ignoreCache?(p(),v(M(c),{key:m.fullPath})):(p(),v(he,{key:1,include:f(e)},[(p(),v(M(c),{key:m.fullPath}))],1032,["include"]))]),_:2},1024)]),_:1})}}}),Ge={key:0,class:"layout-navbar"},Je=I({__name:"default-layout",setup(t){const a=R(!1),e=k(),n=ne(),s=V(),b=j(),c=oe();We(!0);const m="60px",g=i(()=>e.navbar);i(()=>e.menu&&!e.topMenu),i(()=>e.hideMenu),i(()=>e.footer),i(()=>e.menuCollapse?48:e.menuWidth),i(()=>e.menuCollapse);const y=i(()=>({...g.value?{paddingTop:m}:{}}));U(()=>n.role,D=>{D&&!c.accessRouter(b)&&s.push({name:"notFound"})});const x=R(!1);return De("toggleDrawerMenu",()=>{x.value=!x.value}),H(()=>{a.value=!0}),(D,T)=>{const w=ye,u=xe;return p(),v(u,{class:A(["layout",{mobile:f(e).hideMenu}])},{default:r(()=>[f(g)?(p(),B("div",Ge)):F("",!0),o(u,{class:"layout-content",style:ge(f(y))},{default:r(()=>[f(e).tabBar?(p(),v(Ue,{key:0})):F("",!0),o(w,null,{default:r(()=>[o(qe)]),_:1})]),_:1},8,["style"])]),_:1},8,["class"])}}});const et=$(Je,[["__scopeId","data-v-530c78a7"]]);export{et as default};