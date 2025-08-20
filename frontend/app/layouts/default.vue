<template>
  <div class="container">
    <nav>
      <ul>
        <li><strong>Acme Corp</strong></li>
      </ul>
      <ul>
        <li>
          <NuxtLink to="/characters">Characters</NuxtLink>
        </li>
      </ul>
    </nav>
    <slot />
  </div>
</template>
