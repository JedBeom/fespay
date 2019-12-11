<template>
  <div id="app" class="content">
    <div id="nav">
      <router-link to="/" class="button">Home</router-link> |
      <router-link to="/about" class="button is-primary">About</router-link>
    </div>
    <router-view/>
  </div>
</template>