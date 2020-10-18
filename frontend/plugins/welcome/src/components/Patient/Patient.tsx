import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { EntMedicalCare } from '../../api/models/EntMedicalCare'; // import interface MedicalCare
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntDisease } from '../../api/models/EntDisease'; // import interface Disease

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface patient {
  gender: string;
  medicalcare: string;
  employee: string;
  disease: string;
  PatientBirthday: Date;
}

const Patient: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [patient, setPatient] = React.useState<
    Partial<patient>
  >({});

  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [medicalcares, setMedicalCares] = React.useState<EntMedicalCare[]>([]);
  const [employees, setEmployees] = React.useState<EntEmployee[]>([]);
  const [diseases, setDiseases] = React.useState<EntDisease[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getGenders = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };

  const getMedicalCares = async () => {
    const res = await http.listMedicalCare({ limit: 10, offset: 0 });
    setMedicalCares(res);
  };

  const getEmployees = async () => {
    const res = await http.listEmployee({ limit: 10, offset: 0 });
    setEmployees(res);
  };

  const getDiseases = async () => {
    const res = await http.listDisease({ limit: 10, offset: 0 });
    setDiseases(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGenders();
    getMedicalCares();
    getEmployees();
    getDiseases();
  }, []);

  // set data to object playlist_video
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Patient;
    const { value } = event.target;
    setPatient({ ...patient, [name]: value });
    console.log(patient);
  };

  // clear input form
  function clear() {
    setPatient({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/patients';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(patient),
    };

    console.log(patient); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`บันทึกประวัติผู้ป่วย`}>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            
            <Grid item xs={12}>
            <TextField
                  required
                  id="outlined-required"
                  label="เลขประจำตัวผู้ป่วย"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>

            <Grid item xs={12}>
            <TextField
                  required
                  id="outlined-required"
                  label="ชื่อ - นามสกุล"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>
            
            <Grid item xs={10}>
            <TextField
                  required
                  id="outlined-required"
                  label="เลขบัตรประจำตัวประชาชน"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>

            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เพศ</InputLabel>
                <Select
                  name="gender"
                  value={patient.gender || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.genderName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={12}>
            <TextField
                  required
                  id="outlined-required"
                  label="ที่อยู่"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>
            
            <Grid item xs={3}>
            <TextField
                  required
                  id="outlined-required"
                  label="เบอร์โทรศัพท์"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>
            
            <Grid item xs={3}>
              <div className={classes.paper}></div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="วันเดือนปีเกิด"
                  name="วันเดือนปีเกิด"
                  type="date"
                  value={patient.PatientBirthday} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
            <TextField
                  required
                  id="outlined-required"
                  label="อายุ"
                  defaultValue=""
                  variant="outlined"
            />
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>โรคประจำตัว</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>โรคประจำตัว</InputLabel>
                <Select
                  name="disease"
                  value={patient.disease || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.diseaseName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>สิทธิในการรักษา</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>สิทธิในการรักษา</InputLabel>
                <Select
                  name="medicalcare"
                  value={patient.medicalcare || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {medicalcares.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.medicalcareName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>พนักงานผู้บันทึก</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>พนักงานผู้บันทึก</InputLabel>
                <Select
                  name="employee"
                  value={patient.employee || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.employeeName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกข้อมูล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Patient;
