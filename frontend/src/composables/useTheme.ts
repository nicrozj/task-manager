import { ref, watchEffect, computed, type ComputedRef, type Ref } from "vue";

export type Theme = "light" | "dark" | "system";

export function useTheme() {
  const themes = ["light", "dark", "system"];
  const currentTheme = ref<Theme>((localStorage.theme as Theme) || "system");

  const applyTheme = (theme: Theme): void => {
    const root = document.documentElement;
    root.classList.remove("dark", "light");

    if (theme === "dark") {
      root.classList.add("dark");
    } else if (theme === "light") {
      root.classList.add("light");
    } else {
      const isDarkPreferred = window.matchMedia(
        "(prefers-color-scheme: dark)"
      ).matches;
      if (isDarkPreferred) {
        root.classList.add("dark");
      }
    }
  };

  watchEffect(() => {
    const theme = currentTheme.value;

    if (!themes.includes(theme)) {
      currentTheme.value = "system";
      return;
    }

    localStorage.theme = theme;
    applyTheme(theme);
  });

  const setTheme = (theme: Theme): void => {
    currentTheme.value = theme;
  };

  const getSystemTheme = (): "light" | "dark" =>
    window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light";

  const isDark = computed(() => {
    if (currentTheme.value === "dark") return true;
    if (currentTheme.value === "light") return false;
    return window.matchMedia("(prefers-color-scheme: dark)").matches;
  });

  const toggleTheme = (): void => {
    if (isDark.value) {
      setTheme("light");
    } else {
      setTheme("dark");
    }
  };

  return {
    themes,
    currentTheme,
    setTheme,
    isDark,
    toggleTheme,
  };
}
