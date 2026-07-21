---
name: Uzazi App Design System
description: Global visual, component, and motion rules for Uzazi web and Android products.
colors:
  canvas: '#fbf9f4'
  soft: '#f4f1ed'
  white: '#ffffff'
  ink: '#1b1c19'
  copy: '#5d484d'
  muted: '#745860'
  bloom: '#ad2f5b'
  plum: '#72243e'
  plum-soft: '#b75576'
  petal: '#f2cbd5'
  petal-strong: '#f87ea4'
  rose: '#ffd9e2'
typography:
  display:
    fontFamily: 'Charter, Bitstream Charter, Sitka Text, Cambria, serif'
    fontSize: 'clamp(2.75rem, 5vw, 4.5rem)'
    fontWeight: 700
    lineHeight: 1.02
    letterSpacing: '-0.025em'
  headline:
    fontFamily: 'Charter, Bitstream Charter, Sitka Text, Cambria, serif'
    fontSize: 'clamp(2rem, 3vw, 3rem)'
    fontWeight: 600
    lineHeight: 1.08
    letterSpacing: '-0.025em'
  title:
    fontFamily: 'Avenir Next, Avenir, Segoe UI, ui-sans-serif, system-ui, sans-serif'
    fontSize: '1.5rem'
    fontWeight: 500
    lineHeight: 1.2
    letterSpacing: '-0.02em'
  body:
    fontFamily: 'Avenir Next, Avenir, Segoe UI, ui-sans-serif, system-ui, sans-serif'
    fontSize: '1rem'
    fontWeight: 400
    lineHeight: 1.75
  label:
    fontFamily: 'Avenir Next, Avenir, Segoe UI, ui-sans-serif, system-ui, sans-serif'
    fontSize: '0.875rem'
    fontWeight: 500
    lineHeight: 1.4
rounded:
  sm: '0.5rem'
  md: '0.75rem'
  lg: '1rem'
  xl: '1.5rem'
  card: '1rem'
  arch: '4.5rem 4.5rem 10rem 4.5rem / 4.5rem 4.5rem 3rem 4.5rem'
  full: '9999px'
spacing:
  xs: '0.25rem'
  sm: '0.5rem'
  md: '1rem'
  lg: '1.5rem'
  xl: '2rem'
  section-min: '5rem'
  section-max: '8rem'
  shell-gutter-web: '1.5rem'
  shell-gutter-mobile: '1rem'
components:
  button-primary:
    backgroundColor: '{colors.bloom}'
    textColor: '{colors.white}'
    rounded: '{rounded.full}'
    padding: '0.875rem 2rem'
    height: '3rem'
  button-secondary:
    backgroundColor: 'transparent'
    textColor: '{colors.ink}'
    rounded: '{rounded.full}'
    padding: '0.875rem 2rem'
    height: '3rem'
  card-soft:
    backgroundColor: '{colors.soft}'
    textColor: '{colors.ink}'
    rounded: '{rounded.xl}'
    padding: '2rem'
  card-plum:
    backgroundColor: '{colors.plum-soft}'
    textColor: '{colors.white}'
    rounded: '{rounded.xl}'
    padding: '2rem'
  input-default:
    backgroundColor: '{colors.white}'
    textColor: '{colors.ink}'
    rounded: '{rounded.md}'
    padding: '0.75rem 1rem'
    height: '3rem'
---

# Design System: Uzazi App Design System

## 1. Overview

**Creative North Star: “A blooming village companion.”**

Uzazi should feel like a calm village elder, a warm journal, and a gentle wellness companion in one product. The system is global across web, dashboard surfaces, and Android. The Astro marketing pages are the current visual reference implementation, but this document defines the app-wide rules rather than Astro-only implementation details.

The product language is botanical, maternal, culturally rooted, and optimistic. It should avoid hospital utility, generic SaaS dashboards, crypto-style purple gradients, and sterile productivity software. Interfaces should feel soft without becoming childish, supportive without becoming clinical, and emotionally warm without losing clarity.

Use the same design values everywhere, then translate them to each platform. On web, that means Tailwind tokens, fluid layout, hover states, and scroll choreography. On Android, that means Material 3 primitives, 48dp touch targets, edge-to-edge safe areas, and Compose motion that uses the same timing, softness, and restraint.

**Key Characteristics:**

- Botanical growth metaphors: bloom, garden, village, ritual, care, and progress.
- Warm off-white canvas with plum and bloom accents rather than cold medical blues.
- Generous breathing room, short text lines, large touch targets, and calm hierarchy.
- Rounded but not cartoonish shapes: pills for actions, 12-24px/dp for most surfaces.
- Motion that feels like a soft reveal or gentle lift, never bounce, spin, or elastic play.
- Accessibility is part of the brand: readable contrast, reduced motion, focus states, and respectful copy.

## 2. Colors

The Uzazi palette is a floral-plum system: a low-glare botanical canvas, deep plum authority, and bloom pink for growth moments and primary action.

### Primary

- **Bloom Pink** (`#ad2f5b`): The primary action and growth color. Use for main CTAs, active states, focus rings, icons that indicate care or progress, and small moments of celebration. Keep it rare enough to matter.
- **Deep Plum** (`#72243e`): The grounding brand color. Use for high-emphasis sections, footer backgrounds, important headings on colored surfaces, and states that need authority without black harshness.

### Secondary

- **Petal Strong** (`#f87ea4`): A bright botanical accent for selected states, small graphic highlights, badges, and playful but limited emphasis.
- **Plum Soft** (`#b75576`): A softer filled surface for feature cards, Android containers, and illustrations where deep plum would be too heavy.

### Tertiary

- **Rose Wash** (`#ffd9e2`): Gentle supportive fill for icon circles, empty states, mascot bubbles, and quiet calls to action.
- **Petal Mist** (`#f2cbd5`): Borders, dividers, subtle chips, and light outlines. It should structure the interface without making it feel boxed in.

### Neutral

- **Canvas** (`#fbf9f4`): The global page and screen background. Low-glare, warm, and paper-like.
- **Soft Surface** (`#f4f1ed`): Secondary sections, grouped areas, and quiet app surfaces.
- **White** (`#ffffff`): Cards, form fields, navigation menus, and foreground panels.
- **Ink** (`#1b1c19`): Primary text. Use instead of pure black.
- **Copy Plum** (`#5d484d`): Body text on light backgrounds.
- **Muted Plum** (`#745860`): Secondary copy, metadata, helper text, and low-emphasis labels.

### Named Rules

**The Bloom Rarity Rule.** Bloom Pink should usually occupy less than 10% of a screen. If everything blooms, nothing feels like growth.

**The No Hospital Blue Rule.** Do not introduce generic healthcare blues for trust. Trust comes from clarity, privacy copy, spacing, and respectful tone.

**The Plum Contrast Rule.** Use Ink or Deep Plum for meaningful text. Muted Plum is for supporting text only and must still meet contrast requirements.

## 3. Typography

**Display Font:** Charter, Bitstream Charter, Sitka Text, Cambria, serif  
**Body Font:** Avenir Next, Avenir, Segoe UI, system sans-serif  
**Label Font:** Same sans-serif stack as body

**Character:** Display type carries warmth, story, and editorial softness. The sans-serif body stack carries clarity and app usability. Together they should feel human and trustworthy, not decorative or clinical.

### Hierarchy

- **Display** (`700`, `clamp(2.75rem, 5vw, 4.5rem)`, `1.02-1.08`, `-0.025em`): Use for major screen and page heroes. Keep lines short and balanced. On Android, map this to the largest expressive headline style, not long paragraphs.
- **Headline** (`600`, `clamp(2rem, 3vw, 3rem)`, `1.08`, `-0.025em`): Use for major sections, onboarding steps, empty states, and modal titles.
- **Title** (`500`, `1.25-1.5rem`, `1.2`, `-0.02em`): Use for cards, feature groups, settings groups, and form sections.
- **Body** (`400`, `1rem`, `1.65-1.8`): Use for explanatory text, descriptions, and long-form copy. Keep web line lengths near 44-65ch. On mobile, avoid dense paragraphs longer than 3-4 lines.
- **Label** (`500`, `0.8125-0.875rem`, `1.4`): Use for form labels, nav links, metadata, and helper labels. Avoid all-caps except for very short system labels.

### Named Rules

**The Gentle Clarity Rule.** Never lower contrast just to make text feel soft. Softness comes from color temperature, spacing, and tone, not unreadable gray.

**The No Eyebrow Scaffold Rule.** Do not place tiny uppercase tracked labels above every section. One deliberate label is allowed; repeating it becomes template grammar.

**The Short Hero Rule.** Hero headings should be emotionally direct and visually short. If a headline wraps into more than three lines on mobile, rewrite it or reduce the scale.

## 4. Elevation

Uzazi uses tonal layering first and soft shadow second. Most depth should come from warm surface shifts, rounded containers, image masks, and color contrast. Shadows are reserved for interactive lift, hero imagery, floating encouragement cards, and buttons that need tactile affordance.

### Shadow Vocabulary

- **Button Glow** (`0 6px 12px rgb(165 46 90 / .16)`): Primary buttons and single dominant actions. It should feel like a soft bloom, not a drop shadow.
- **Hero Lift** (`0 18px 45px rgb(114 36 62 / .12)`): Large image masks, hero arches, and major visual anchors only.
- **Companion Card Lift** (`0 10px 24px rgb(90 45 60 / .13)`): Mascot prompts, encouragement cards, or transient support surfaces.
- **Hover Lift** (`0 7px 8px rgb(114 36 62 / .08)`): Feature/card hover response on web. On Android, translate as tonal elevation plus a small scale or color-state change instead of heavy shadow.

### Named Rules

**The Flat-Until-Alive Rule.** Surfaces are calm at rest. Elevation appears when a component is interactive, floating over content, or narratively important.

**The No Ghost Card Rule.** Do not combine a 1px border with a large generic shadow on ordinary cards. Pick tonal layering, a thin Petal border, or a purposeful soft lift.

**The Soft Edge Rule.** Default cards and input fields sit around 12-24px/dp radius. Do not use 32px+ radii on ordinary cards; reserve full pills for buttons, chips, and progress forms.

## 5. Components

Components should feel tactile, generous, and emotionally safe. Platform-native behavior matters: web components can use hover and scroll reveals; Android components should use Material interaction states, haptics where appropriate, and Compose animations that match the same motion values.

### Buttons

- **Shape:** Full pill for primary and secondary actions. Minimum height is 48px on web and 48dp on Android.
- **Primary:** Bloom Pink fill, White text, Button Glow shadow, medium/semibold label. Hover on web moves `translateY(-2px)` or less and darkens to Deep Plum. Press scales to about `0.97-0.98`.
- **Secondary:** Transparent or White background, Petal border, Ink text. Hover/focus may shift to White or Soft Surface.
- **Focus:** Always visible. Web uses a 2px Bloom outline with 4px offset. Android uses the platform focus/pressed indication tinted Bloom or Rose.
- **Motion:** Transition transform, background, and shadow over `220-300ms` using `cubic-bezier(.22, 1, .36, 1)`. Never bounce.

### Cards / Containers

- **Feature Cards:** Large rounded containers, usually `rounded-2xl` or 16-24dp. Use asymmetric grid spans when screen size allows. Avoid endless identical icon-card grids.
- **Support Cards:** White or Rose surfaces with 16-24px/dp radius and 20-40px/dp padding. Use for Mama Bear prompts, privacy reassurance, or contextual help.
- **High-Emphasis Cards:** Plum Soft or Rose fills with high-contrast text. Keep copy short.
- **Image Containers:** Use distinctive organic masks, such as the hero arch, for brand moments. On Android, translate the arch into rounded top corners plus one exaggerated bottom corner only where technically and visually appropriate.
- **Internal Padding:** 24-40px on web marketing cards; 16-24dp on Android app cards.

### Inputs / Fields

- **Style:** White fill, Petal border, 12px/dp radius, 48px/dp minimum height, Ink text, Muted Plum placeholder.
- **Focus:** Border shifts to Bloom and receives a soft Bloom ring or platform focus indication.
- **Privacy Copy:** When asking for sensitive journey information, include concise reassurance near the form. Do not ask for private medical details unless the feature truly requires it.
- **Errors:** Use clear text plus an error color. Do not rely on color alone. Preserve the soft tone while being direct.

### Navigation

- **Header/App Bars:** Light surfaces use Canvas with subtle Petal divider and, on web, a gentle backdrop blur. Android top app bars should use Canvas or White with tonal separation rather than heavy shadow.
- **Links:** Copy Plum by default, Bloom on hover/active. Web nav underline animates with scaleX over `220ms` from right-to-left on exit and left-to-right on enter.
- **Mobile Menus:** Use full-width, fixed, or platform-native sheets. Do not render menus inside clipped containers.
- **Bottom Navigation:** Android can use Material bottom navigation if needed, tinted with Ink/Muted/Bloom and a Rose active container.

### Iconography and Mascot Language

- **Icons:** Rounded, botanical, and simple. Use line or solid icons with clear silhouettes. Avoid generic corporate outline icon sets when a custom botanical or care symbol would communicate better.
- **Icon Tiles:** 48px/dp rounded squares or circles with Rose/Bloom tonal fills. On hover, web icon tiles may lift `-2px`, rotate `-3deg`, and scale to `1.06`.
- **Mama Bear / Companion Prompts:** The companion voice is encouraging and brief. Use sparingly for reassurance, onboarding, streak recovery, and celebration. The mascot should never block tasks or trivialize serious health content.

### Progress and Wellness Patterns

- **Garden Progress:** Prefer growth metaphors over generic progress bars when tracking wellness, rituals, or mood consistency.
- **Ritual Check-ins:** Keep flows short, one primary action per step, with gentle confirmation after completion.
- **Community / Village Patterns:** Use grouped cards, human quotes, local context, and safety/privacy language. Community UI must feel moderated and respectful.
- **Celebration:** Use small Bloom/Rose animations, icon changes, or copy moments. Avoid confetti overload.

### Motion and Animation Patterns

- **Global Ease:** Use `cubic-bezier(.22, 1, .36, 1)` as the default expressive ease-out. It should feel soft and settled.
- **Hero Copy Entrance:** Fade from `opacity: 0`, `translateY(18px)`, `blur(4px)` to rest over `620ms`, stagger children by `70ms`.
- **Hero Visual Entrance:** Fade and slide from `translateX(22px) scale(.985)` to rest over `760ms`, delayed about `90ms` after copy starts.
- **Scroll Reveal:** Content is visible by default. On capable clients, animate from `opacity: .82`, `translateY(14px)`, `blur(2px)` to rest over `520ms`, with 70ms item staggers. Never gate visibility on JavaScript.
- **Menu Entrance:** Mobile menu/sheet appears with a small `translateY(-6px) scale(.97)` to rest over `180ms`, origin at top/right or sheet edge.
- **Micro Lift:** Buttons and cards lift only a few pixels. Feature cards can lift up to `6px`; buttons should stay around `2px`.
- **Android Translation:** Use Compose `tween(durationMillis = 180-760, easing = CubicBezierEasing(0.22f, 1f, 0.36f, 1f))`. Use critically damped springs only for platform-native gestures; do not add bounce or elastic effects.
- **Reduced Motion:** Required on every platform. Web uses `prefers-reduced-motion`; Android respects system animator duration scale and reduced motion/accessibility settings. Replace movement with instant state changes or gentle crossfades.

## 6. Do's and Don'ts

### Do:

- **Do** treat this as the global app design system. Astro, Next.js, and Android should all translate these rules through their platform conventions.
- **Do** use Canvas, Soft Surface, White, Ink, Copy Plum, Bloom Pink, Deep Plum, Rose, and Petal as the default palette before adding new colors.
- **Do** keep primary actions obvious, rare, and Bloom Pink.
- **Do** make every tappable/clickable target at least 44px on web and 48dp on Android.
- **Do** use the garden, village, ritual, and Mama Bear patterns to make wellness progress feel emotionally supportive.
- **Do** keep sensitive flows calm, private, and clear. Use reassurance copy and avoid unnecessary data collection.
- **Do** use real imagery or meaningful product visuals for brand surfaces. The current web reference uses warm mother-and-baby and growth-ritual imagery.
- **Do** include visible focus states, semantic labels, accessible contrast, and reduced-motion alternatives.
- **Do** let Android use Material 3 structure while preserving Uzazi color, radius, typography, and motion character.
- **Do** keep motion purposeful: page entrance, sheet/menu entry, micro-lift, garden growth, completion feedback, and small mascot moments.

### Don't:

- **Don't** make this Astro-specific. Do not document class names as the source of truth; document the underlying design rule and token.
- **Don't** use hospital blue, generic wellness teal, neon purple gradients, or cold grayscale dashboards as the Uzazi default.
- **Don't** use repeated tiny uppercase eyebrows, numbered section markers, or identical icon-card grids as page scaffolding.
- **Don't** rely on emoji as the primary visual system. Emoji can support journaling moments, but icons, copy, layout, and mascot language carry the brand.
- **Don't** animate layout properties or create bouncy/elastic motion. Uzazi motion settles softly.
- **Don't** hide content until animation JavaScript runs. Content must be readable if animation fails.
- **Don't** use heavy shadows on ordinary cards or pair borders with large generic shadows.
- **Don't** over-round normal cards past 24px/dp. Save full pills for actions, chips, and progress rails.
- **Don't** make Mama Bear cute at the expense of trust. Serious health, safety, or privacy content needs respectful clarity first.
- **Don't** ship Android screens that feel like a web page squeezed into a phone. Translate spacing, navigation, motion, and touch ergonomics into native patterns.
