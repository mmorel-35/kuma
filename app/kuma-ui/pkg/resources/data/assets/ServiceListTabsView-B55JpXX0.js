import{d as R,s as C,V as L,r as s,o as a,m as v,w as t,k as _,e as n,b as c,c as h,L as w,M as g,n as B,t as M,p as N}from"./index-Di1IipgC.js";const S={class:"stack"},T=["innerHTML"],y=R({__name:"ServiceListTabsView",props:{mesh:{}},setup(V){const r=C(),l=V;return L(()=>r.currentRoute.value.name,m=>{var i;m==="service-list-tabs-view"&&r.replace(((i=l.mesh.meshServices)==null?void 0:i.enabled)==="Exclusive"?{name:"mesh-service-list-view"}:{name:"service-list-view"})},{immediate:!0}),(m,i)=>{const f=s("XAction"),b=s("XActionGroup"),x=s("RouterView"),k=s("AppView"),A=s("RouteView");return a(),v(A,{name:"service-list-tabs-view",params:{mesh:""}},{default:t(({route:o,t:u})=>[_("div",S,[_("div",{innerHTML:u("services.routes.items.intro",{},{defaultMessage:""})},null,8,T),n(),c(k,null,{actions:t(()=>[c(b,{expanded:!0},{default:t(()=>[(a(!0),h(w,null,g(o.children,({name:e})=>{var d,p;return a(),h(w,{key:e},[((d=l.mesh.meshServices)==null?void 0:d.enabled)==="Exclusive"&&["service-list-view","external-service-list-view"].includes(e)?N("",!0):(a(),v(f,{key:0,class:B({active:((p=o.child())==null?void 0:p.name)===e}),to:{name:e,params:{mesh:o.params.mesh}},"data-testid":`${e}-sub-tab`},{default:t(()=>[n(M(u(`services.routes.items.navigation.${e}`)),1)]),_:2},1032,["class","to","data-testid"]))],64)}),128))]),_:2},1024)]),default:t(()=>[n(),c(x)]),_:2},1024)])]),_:1})}}});export{y as default};