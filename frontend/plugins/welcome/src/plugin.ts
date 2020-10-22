import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Patient from './components/Patient';
import Signin from './components/Signin';
import MenuTab from './components/Menu';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Signin);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/patients', Patient);
    router.registerRoute('/MenuTab', MenuTab);

  }
});
