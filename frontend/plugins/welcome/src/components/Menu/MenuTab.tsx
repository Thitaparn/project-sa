import React, { FC } from 'react';
//import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
 // Link,
} from '@backstage/core';
import { emphasize, withStyles, Theme } from '@material-ui/core/styles';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Chip from '@material-ui/core/Chip';
import HomeIcon from '@material-ui/icons/Home';
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
const MenuTab: FC<{}> = () => {
    const StyledBreadcrumb = withStyles((theme: Theme) => ({
        root: {
          backgroundColor: theme.palette.grey[100],
          height: theme.spacing(3),
          color: theme.palette.grey[800],
          fontWeight: theme.typography.fontWeightRegular,
          '&:hover, &:focus': {
            backgroundColor: theme.palette.grey[300],
          },
          '&:active': {
            boxShadow: theme.shadows[1],
            backgroundColor: emphasize(theme.palette.grey[300], 0.12),
          },
        },
      }))(Chip) as typeof Chip; // TypeScript only: need a type cast here because https://github.com/Microsoft/TypeScript/issues/26591
      function handleClick(event: React.MouseEvent<Element, MouseEvent>) {
        event.preventDefault();
        console.info('You clicked a breadcrumb.');
      }

  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ระบบบันทึกประวัติทันตกรรม">
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
      </Header>
      <Content>
        <ContentHeader title="Menu"></ContentHeader>
        <Breadcrumbs aria-label="breadcrumb">
            <StyledBreadcrumb
               component="a"
               href="#"
               label="Home"
               icon={<HomeIcon fontSize="large" />}
               onClick={handleClick}
            />
        <StyledBreadcrumb component="a" href="#" label="Catalog" onClick={handleClick} />
            <StyledBreadcrumb
                label="Accessories"
                deleteIcon={<ExpandMoreIcon />}
                onClick={handleClick}
                onDelete={handleClick}
            />
            </Breadcrumbs>
      </Content>
    </Page>
  );
};
export default MenuTab;