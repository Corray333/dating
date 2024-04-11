<template>
    <section class="flex w-full relative justify-center p-5">
        <section class="relative wm w-full">
            <navbar class="absolute -left-60"/>
            <main class="p-5 w-full bg-dark h-fit rounded-xl flex flex-col items-center shadow-lg">
                <slot />
            </main>
        </section>
    </section>
</template>