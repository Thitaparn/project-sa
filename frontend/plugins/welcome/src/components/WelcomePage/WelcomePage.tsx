import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import { Grid } from '@material-ui/core';

const WelcomePage: FC<{}> = () => {
  const profile = { givenName: 'Dentals System' };

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Welcome ${profile.givenName || 'to Backstage'}`}
      ></Header>
      <Content>
        <ContentHeader title="ประวัติผู้ป่วย">
          
        <Link component={RouterLink} to="/">
          <Button variant="contained" color="primary" disableElevation>
                ComeBack Sign In
          </Button>
          </Link>
          
          <Link component={RouterLink} to="/patients">
          <Button variant="contained" color="primary" disableElevation>
                Add Patient
          </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
      </Content>
    </Page>
  );
};

export default WelcomePage;
