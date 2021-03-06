package goscore_test

import (
	"testing"
	"github.com/asafschers/goscore"
	"encoding/xml"
	"fmt"
)

const randomForestXml = `
<?xml version="1.0"?>
<PMML version="4.3" xmlns="http://www.dmg.org/PMML-4_3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.dmg.org/PMML-4_3 http://www.dmg.org/pmml/v4-3/pmml-4-3.xsd">
 <Header copyright="Copyright (c) 2017 asafschers1" description="Random Forest Tree Model">
  <Extension name="user" value="asafschers1" extender="Rattle/PMML"/>
  <Application name="Rattle/PMML" version="1.4"/>
  <Timestamp>2017-05-18 15:39:02</Timestamp>
 </Header>
 <DataDictionary numberOfFields="12">
  <DataField name="Survived" optype="categorical" dataType="string">
   <Value value="0"/>
   <Value value="1"/>
  </DataField>
  <DataField name="PassengerId" optype="continuous" dataType="double"/>
  <DataField name="Pclass" optype="continuous" dataType="double"/>
  <DataField name="Name" optype="categorical" dataType="string"/>
  <DataField name="Sex" optype="categorical" dataType="string"/>
  <DataField name="Age" optype="continuous" dataType="double"/>
  <DataField name="SibSp" optype="continuous" dataType="double"/>
  <DataField name="Parch" optype="continuous" dataType="double"/>
  <DataField name="Ticket" optype="categorical" dataType="string"/>
  <DataField name="Fare" optype="continuous" dataType="double"/>
  <DataField name="Cabin" optype="categorical" dataType="string"/>
  <DataField name="Embarked" optype="categorical" dataType="string"/>
 </DataDictionary>
 <MiningModel modelName="randomForest_Model" algorithmName="randomForest" functionName="classification">
  <MiningSchema>
   <MiningField name="Survived" usageType="predicted"/>
   <MiningField name="PassengerId" usageType="active"/>
   <MiningField name="Pclass" usageType="active"/>
   <MiningField name="Name" usageType="active"/>
   <MiningField name="Sex" usageType="active"/>
   <MiningField name="Age" usageType="active"/>
   <MiningField name="SibSp" usageType="active"/>
   <MiningField name="Parch" usageType="active"/>
   <MiningField name="Ticket" usageType="active"/>
   <MiningField name="Fare" usageType="active"/>
   <MiningField name="Cabin" usageType="active"/>
   <MiningField name="Embarked" usageType="active"/>
  </MiningSchema>
  <Output>
   <OutputField name="Predicted_Survived" feature="predictedValue"/>
   <OutputField name="Probability_0" optype="continuous" dataType="double" feature="probability" value="0"/>
   <OutputField name="Probability_1" optype="continuous" dataType="double" feature="probability" value="1"/>
  </Output>
  <Segmentation multipleModelMethod="majorityVote">
   <Segment id="1">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Fare" operator="lessOrEqual" value="51.9312499999999986"/>
       <Node id="4">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;female&quot;</Array>
        </SimpleSetPredicate>
        <Node id="8">
         <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
         <Node id="14">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
          <Node id="26">
           <SimplePredicate field="Age" operator="lessOrEqual" value="56"/>
           <Node id="44">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="70" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="49"/>
            </Node>
            <Node id="71" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="49"/>
            </Node>
           </Node>
           <Node id="45">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="72">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="106">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="130" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;Q&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="131">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="154" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
               </Node>
               <Node id="155" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="37"/>
               </Node>
              </Node>
             </Node>
             <Node id="107" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="73">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="108" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
             </Node>
             <Node id="109" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="27">
           <SimplePredicate field="Age" operator="greaterThan" value="56"/>
           <Node id="46" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="57.5"/>
           </Node>
           <Node id="47" score="1">
            <SimplePredicate field="Age" operator="greaterThan" value="57.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="15" score="1">
          <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
         <Node id="16">
          <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
          <Node id="28">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="21.7041499999999985"/>
           <Node id="48">
            <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
            <Node id="74">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="110">
              <SimplePredicate field="Age" operator="lessOrEqual" value="14.5"/>
              <Node id="132">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="9.54795000000000016"/>
               <Node id="156" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="13.5"/>
               </Node>
               <Node id="157" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="13.5"/>
               </Node>
              </Node>
              <Node id="133" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="9.54795000000000016"/>
              </Node>
             </Node>
             <Node id="111" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="14.5"/>
             </Node>
            </Node>
            <Node id="75">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="112">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              <Node id="134" score="1">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="135" score="1">
               <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
             <Node id="113" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="49">
            <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
            <Node id="76">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="114" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="115">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
              <Node id="136" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7"/>
              </Node>
              <Node id="137">
               <SimplePredicate field="Fare" operator="greaterThan" value="7"/>
               <Node id="158" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.70000000000000018"/>
               </Node>
               <Node id="159">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.70000000000000018"/>
                <Node id="172">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="23.5"/>
                 <Node id="186" score="1">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;Q&quot;</Array>
                  </SimpleSetPredicate>
                 </Node>
                 <Node id="187" score="1">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                  </SimpleSetPredicate>
                 </Node>
                </Node>
                <Node id="173">
                 <SimplePredicate field="Age" operator="greaterThan" value="23.5"/>
                 <Node id="188" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
                 </Node>
                 <Node id="189">
                  <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
                  <Node id="196" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
                  </Node>
                  <Node id="197" score="1">
                   <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
                  </Node>
                 </Node>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="77" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="29" score="0">
           <SimplePredicate field="Fare" operator="greaterThan" value="21.7041499999999985"/>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
          <Node id="30">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
           </SimpleSetPredicate>
           <Node id="50">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="15.3728999999999996"/>
            <Node id="78" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="8.00835000000000008"/>
            </Node>
            <Node id="79" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="8.00835000000000008"/>
            </Node>
           </Node>
           <Node id="51">
            <SimplePredicate field="Fare" operator="greaterThan" value="15.3728999999999996"/>
            <Node id="80" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="18.9291499999999999"/>
            </Node>
            <Node id="81" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="18.9291499999999999"/>
            </Node>
           </Node>
          </Node>
          <Node id="31">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="52">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
            <Node id="82">
             <SimplePredicate field="Age" operator="lessOrEqual" value="54"/>
             <Node id="116" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="15.25"/>
             </Node>
             <Node id="117">
              <SimplePredicate field="Fare" operator="greaterThan" value="15.25"/>
              <Node id="138" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="17.6999999999999993"/>
              </Node>
              <Node id="139">
               <SimplePredicate field="Fare" operator="greaterThan" value="17.6999999999999993"/>
               <Node id="160">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                <Node id="174" score="0">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                </Node>
                <Node id="175" score="1">
                 <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
               <Node id="161">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                <Node id="176" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
                </Node>
                <Node id="177">
                 <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
                 <Node id="190" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
                 </Node>
                 <Node id="191" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
                 </Node>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="83" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="54"/>
            </Node>
           </Node>
           <Node id="53">
            <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
            <Node id="84" score="1">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;male&quot;</Array>
        </SimpleSetPredicate>
        <Node id="10">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="1" type="string">&quot;C&quot;</Array>
         </SimpleSetPredicate>
         <Node id="18">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="15.1478999999999999"/>
          <Node id="32" score="0">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
          </Node>
          <Node id="33">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           <Node id="54">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="8.58960000000000079"/>
            <Node id="86">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="118">
              <SimplePredicate field="Age" operator="lessOrEqual" value="22.75"/>
              <Node id="140" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="5.61875000000000036"/>
              </Node>
              <Node id="141" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="5.61875000000000036"/>
              </Node>
             </Node>
             <Node id="119">
              <SimplePredicate field="Age" operator="greaterThan" value="22.75"/>
              <Node id="142" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.22710000000000008"/>
              </Node>
              <Node id="143" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.22710000000000008"/>
              </Node>
             </Node>
            </Node>
            <Node id="87" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="8.58960000000000079"/>
           </Node>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Fare" operator="greaterThan" value="15.1478999999999999"/>
          <Node id="34">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           <Node id="56">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            <Node id="88" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="57"/>
            </Node>
            <Node id="89" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="57"/>
            </Node>
           </Node>
           <Node id="57">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            <Node id="90" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="33.5374999999999943"/>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="33.5374999999999943"/>
            </Node>
           </Node>
          </Node>
          <Node id="35">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           <Node id="58" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="15.5"/>
           </Node>
           <Node id="59">
            <SimplePredicate field="Age" operator="greaterThan" value="15.5"/>
            <Node id="92">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="120" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="20.2333500000000015"/>
             </Node>
             <Node id="121" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="20.2333500000000015"/>
             </Node>
            </Node>
            <Node id="93" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20">
          <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
          <Node id="36" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="11"/>
          </Node>
          <Node id="37">
           <SimplePredicate field="Age" operator="greaterThan" value="11"/>
           <Node id="60">
            <SimplePredicate field="Age" operator="lessOrEqual" value="63"/>
            <Node id="94" score="0">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;Q&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="95">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="122">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="144">
               <SimplePredicate field="Age" operator="lessOrEqual" value="61.5"/>
               <Node id="162">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                <Node id="178">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="38.9500000000000028"/>
                 <Node id="192" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="38"/>
                 </Node>
                 <Node id="193">
                  <SimplePredicate field="Age" operator="greaterThan" value="38"/>
                  <Node id="198" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="48.25"/>
                  </Node>
                  <Node id="199">
                   <SimplePredicate field="Age" operator="greaterThan" value="48.25"/>
                   <Node id="202" score="1">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="29.4354000000000013"/>
                   </Node>
                   <Node id="203" score="0">
                    <SimplePredicate field="Fare" operator="greaterThan" value="29.4354000000000013"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="179" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="38.9500000000000028"/>
                </Node>
               </Node>
               <Node id="163" score="0">
                <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
              <Node id="145" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="61.5"/>
              </Node>
             </Node>
             <Node id="123" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="63"/>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
          <Node id="38">
           <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
           <Node id="62" score="1">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="63">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            <Node id="96" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
            </Node>
            <Node id="97" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
            </Node>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
           <Node id="64">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="98" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="99" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
           <Node id="65">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            <Node id="100" score="0">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;Q&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="101">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="124">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              <Node id="146" score="0">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
              </Node>
              <Node id="147">
               <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
               <Node id="164" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="18.0749999999999993"/>
               </Node>
               <Node id="165" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="18.0749999999999993"/>
               </Node>
              </Node>
             </Node>
             <Node id="125" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Fare" operator="greaterThan" value="51.9312499999999986"/>
       <Node id="6">
        <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
        <Node id="12">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22" score="1">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
         </Node>
         <Node id="23">
          <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
          <Node id="40" score="1">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="135.775000000000006"/>
          </Node>
          <Node id="41">
           <SimplePredicate field="Fare" operator="greaterThan" value="135.775000000000006"/>
           <Node id="66" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="23"/>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="23"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="24">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="387.664600000000007"/>
          <Node id="42" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
          </Node>
          <Node id="43">
           <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
           <Node id="68">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="59.0874999999999986"/>
            <Node id="102" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="22"/>
            </Node>
            <Node id="103">
             <SimplePredicate field="Age" operator="greaterThan" value="22"/>
             <Node id="126">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="148" score="0">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="2"/>
              </Node>
              <Node id="149" score="1">
               <SimplePredicate field="Pclass" operator="greaterThan" value="2"/>
              </Node>
             </Node>
             <Node id="127">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              <Node id="150">
               <SimplePredicate field="Age" operator="lessOrEqual" value="49.5"/>
               <Node id="166" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="52.8271000000000015"/>
               </Node>
               <Node id="167" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="52.8271000000000015"/>
               </Node>
              </Node>
              <Node id="151" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="49.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="69">
            <SimplePredicate field="Fare" operator="greaterThan" value="59.0874999999999986"/>
            <Node id="104" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="105">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="128">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              <Node id="152">
               <SimplePredicate field="Age" operator="lessOrEqual" value="48.5"/>
               <Node id="168">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
                <Node id="180" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="99.9895999999999958"/>
                </Node>
                <Node id="181" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="99.9895999999999958"/>
                </Node>
               </Node>
               <Node id="169">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="182" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="86.7374999999999972"/>
                </Node>
                <Node id="183" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="86.7374999999999972"/>
                </Node>
               </Node>
              </Node>
              <Node id="153">
               <SimplePredicate field="Age" operator="greaterThan" value="48.5"/>
               <Node id="170" score="0">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="171">
                <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                <Node id="184">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
                 <Node id="194">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;C&quot;</Array>
                  </SimpleSetPredicate>
                  <Node id="200" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="54.5"/>
                  </Node>
                  <Node id="201" score="1">
                   <SimplePredicate field="Age" operator="greaterThan" value="54.5"/>
                  </Node>
                 </Node>
                 <Node id="195" score="0">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                  </SimpleSetPredicate>
                 </Node>
                </Node>
                <Node id="185" score="0">
                 <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="129" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="25" score="1">
          <SimplePredicate field="Fare" operator="greaterThan" value="387.664600000000007"/>
         </Node>
        </Node>
       </Node>
       <Node id="7" score="0">
        <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="2">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
       <Node id="4">
        <SimplePredicate field="Fare" operator="lessOrEqual" value="12.3125"/>
        <Node id="8" score="0">
         <SimplePredicate field="Age" operator="lessOrEqual" value="45"/>
        </Node>
        <Node id="9" score="0">
         <SimplePredicate field="Age" operator="greaterThan" value="45"/>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Fare" operator="greaterThan" value="12.3125"/>
        <Node id="10">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="54.0499999999999972"/>
         <Node id="16">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="28">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
           <Node id="46" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="12.4375"/>
           </Node>
           <Node id="47">
            <SimplePredicate field="Fare" operator="greaterThan" value="12.4375"/>
            <Node id="72">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="30.5978999999999992"/>
             <Node id="104">
              <SimplePredicate field="Age" operator="lessOrEqual" value="49.5"/>
              <Node id="128">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="148">
                <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
                <Node id="166" score="1">
                 <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;female&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
                <Node id="167">
                 <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;male&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="182" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                 </Node>
                 <Node id="183" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                 </Node>
                </Node>
               </Node>
               <Node id="149">
                <SimplePredicate field="Age" operator="greaterThan" value="32"/>
                <Node id="168" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="25.7583500000000001"/>
                </Node>
                <Node id="169" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="25.7583500000000001"/>
                </Node>
               </Node>
              </Node>
              <Node id="129">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               <Node id="150">
                <SimplePredicate field="Age" operator="lessOrEqual" value="39.5"/>
                <Node id="170" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="12.7624999999999993"/>
                </Node>
                <Node id="171">
                 <SimplePredicate field="Fare" operator="greaterThan" value="12.7624999999999993"/>
                 <Node id="184" score="1">
                  <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;female&quot;</Array>
                  </SimpleSetPredicate>
                 </Node>
                 <Node id="185">
                  <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;male&quot;</Array>
                  </SimpleSetPredicate>
                  <Node id="194" score="0">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;C&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="195" score="0">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="151">
                <SimplePredicate field="Age" operator="greaterThan" value="39.5"/>
                <Node id="172">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="43"/>
                 <Node id="186" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="14.375"/>
                 </Node>
                 <Node id="187" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="14.375"/>
                 </Node>
                </Node>
                <Node id="173">
                 <SimplePredicate field="Age" operator="greaterThan" value="43"/>
                 <Node id="188">
                  <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                  <Node id="196" score="1">
                   <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;female&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="197" score="0">
                   <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;male&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                 </Node>
                 <Node id="189" score="0">
                  <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 </Node>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="105">
              <SimplePredicate field="Age" operator="greaterThan" value="49.5"/>
              <Node id="130">
               <SimplePredicate field="Age" operator="lessOrEqual" value="72.5"/>
               <Node id="152" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="51.5"/>
               </Node>
               <Node id="153">
                <SimplePredicate field="Age" operator="greaterThan" value="51.5"/>
                <Node id="174" score="0">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;C&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
                <Node id="175">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="190">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
                  <Node id="198" score="1">
                   <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                  </Node>
                  <Node id="199" score="0">
                   <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
                  </Node>
                 </Node>
                 <Node id="191">
                  <SimplePredicate field="Age" operator="greaterThan" value="53"/>
                  <Node id="200">
                   <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                   <Node id="204" score="1">
                    <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                     <Array n="1" type="string">&quot;female&quot;</Array>
                    </SimpleSetPredicate>
                   </Node>
                   <Node id="205" score="0">
                    <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                     <Array n="1" type="string">&quot;male&quot;</Array>
                    </SimpleSetPredicate>
                   </Node>
                  </Node>
                  <Node id="201">
                   <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
                   <Node id="206">
                    <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                    <Node id="208" score="1">
                     <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                      <Array n="1" type="string">&quot;female&quot;</Array>
                     </SimpleSetPredicate>
                    </Node>
                    <Node id="209" score="0">
                     <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                      <Array n="1" type="string">&quot;male&quot;</Array>
                     </SimpleSetPredicate>
                    </Node>
                   </Node>
                   <Node id="207" score="0">
                    <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="131" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="72.5"/>
              </Node>
             </Node>
            </Node>
            <Node id="73">
             <SimplePredicate field="Fare" operator="greaterThan" value="30.5978999999999992"/>
             <Node id="106">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="132">
               <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
               <Node id="154" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="155">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
                <Node id="176" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="37.5499999999999972"/>
                </Node>
                <Node id="177" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="37.5499999999999972"/>
                </Node>
               </Node>
              </Node>
              <Node id="133" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
              </Node>
             </Node>
             <Node id="107">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              <Node id="134">
               <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
               <Node id="156" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="23"/>
               </Node>
               <Node id="157" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="23"/>
               </Node>
              </Node>
              <Node id="135">
               <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
               <Node id="158" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="159">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
                <Node id="178" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="52.2771000000000043"/>
                </Node>
                <Node id="179" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="52.2771000000000043"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="29" score="1">
           <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="30">
           <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
           <Node id="48" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="49">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="74" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="75" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="31">
           <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
           <Node id="50">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
            <Node id="76">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             <Node id="108">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="29.25"/>
              <Node id="136" score="1">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="137" score="0">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="109" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="29.25"/>
             </Node>
            </Node>
            <Node id="77" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
           <Node id="51">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
            <Node id="78" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="79" score="0">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimplePredicate field="Fare" operator="greaterThan" value="54.0499999999999972"/>
         <Node id="18">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
          <Node id="32">
           <SimplePredicate field="Age" operator="lessOrEqual" value="43.5"/>
           <Node id="52">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="80" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="77.9646000000000043"/>
            </Node>
            <Node id="81">
             <SimplePredicate field="Fare" operator="greaterThan" value="77.9646000000000043"/>
             <Node id="110">
              <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
              <Node id="138" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="23.5"/>
              </Node>
              <Node id="139" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="23.5"/>
              </Node>
             </Node>
             <Node id="111" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="53">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="82">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
             <Node id="112">
              <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
              <Node id="140" score="1">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="141">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="160" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
               </Node>
               <Node id="161" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="113" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
             </Node>
            </Node>
            <Node id="83">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="114" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="115" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="33">
           <SimplePredicate field="Age" operator="greaterThan" value="43.5"/>
           <Node id="54">
            <SimplePredicate field="Age" operator="lessOrEqual" value="63.5"/>
            <Node id="84">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="123.462500000000006"/>
             <Node id="116">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="78.7333500000000015"/>
              <Node id="142">
               <SimplePredicate field="Age" operator="lessOrEqual" value="57"/>
               <Node id="162" score="1">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="163" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
              <Node id="143" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="57"/>
              </Node>
             </Node>
             <Node id="117" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="78.7333500000000015"/>
             </Node>
            </Node>
            <Node id="85" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="123.462500000000006"/>
            </Node>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="63.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          <Node id="34" score="0">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
          </Node>
          <Node id="35" score="1">
           <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
       <Node id="6">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;female&quot;</Array>
        </SimpleSetPredicate>
        <Node id="12">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20">
          <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
          <Node id="36" score="1">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.68125000000000036"/>
          </Node>
          <Node id="37">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.68125000000000036"/>
           <Node id="56" score="1">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="57">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="86">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="8.08334999999999937"/>
             <Node id="118" score="1">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
             </Node>
             <Node id="119" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
             </Node>
            </Node>
            <Node id="87">
             <SimplePredicate field="Fare" operator="greaterThan" value="8.08334999999999937"/>
             <Node id="120" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="11.8187499999999996"/>
             </Node>
             <Node id="121" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="11.8187499999999996"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="21" score="0">
          <SimplePredicate field="Age" operator="greaterThan" value="30"/>
         </Node>
        </Node>
        <Node id="13">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22" score="1">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="17.6000000000000014"/>
         </Node>
         <Node id="23">
          <SimplePredicate field="Fare" operator="greaterThan" value="17.6000000000000014"/>
          <Node id="38">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="22.7374999999999972"/>
           <Node id="58" score="0">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="59">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="88" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="21.5499999999999972"/>
            </Node>
            <Node id="89" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="21.5499999999999972"/>
            </Node>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Fare" operator="greaterThan" value="22.7374999999999972"/>
           <Node id="60" score="0">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="4.5"/>
           </Node>
           <Node id="61">
            <SimplePredicate field="Parch" operator="greaterThan" value="4.5"/>
            <Node id="90" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="39.5"/>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="39.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;male&quot;</Array>
        </SimpleSetPredicate>
        <Node id="14">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="7.9104000000000001"/>
         <Node id="24" score="0">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="7.17499999999999982"/>
         </Node>
         <Node id="25">
          <SimplePredicate field="Fare" operator="greaterThan" value="7.17499999999999982"/>
          <Node id="40">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.23959999999999937"/>
           <Node id="62" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="7.22710000000000008"/>
           </Node>
           <Node id="63">
            <SimplePredicate field="Fare" operator="greaterThan" value="7.22710000000000008"/>
            <Node id="92" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="93" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="41">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.23959999999999937"/>
           <Node id="64">
            <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
            <Node id="94" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
            </Node>
            <Node id="95" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
            </Node>
           </Node>
           <Node id="65">
            <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
            <Node id="96">
             <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
             <Node id="122" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="123" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
            <Node id="97" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="30"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15">
         <SimplePredicate field="Fare" operator="greaterThan" value="7.9104000000000001"/>
         <Node id="26">
          <SimplePredicate field="Age" operator="lessOrEqual" value="14"/>
          <Node id="42" score="1">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
          </Node>
          <Node id="43">
           <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
           <Node id="66" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
           </Node>
           <Node id="67">
            <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
            <Node id="98" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
            </Node>
            <Node id="99" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="27">
          <SimplePredicate field="Age" operator="greaterThan" value="14"/>
          <Node id="44">
           <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
           <Node id="68">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="100">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="48.0916500000000013"/>
             <Node id="124" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="19.5"/>
             </Node>
             <Node id="125">
              <SimplePredicate field="Age" operator="greaterThan" value="19.5"/>
              <Node id="144" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
              </Node>
              <Node id="145">
               <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
               <Node id="164">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
                </SimpleSetPredicate>
                <Node id="180" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                </Node>
                <Node id="181">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 <Node id="192" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="15.3728999999999996"/>
                 </Node>
                 <Node id="193">
                  <SimplePredicate field="Fare" operator="greaterThan" value="15.3728999999999996"/>
                  <Node id="202" score="0">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;C&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="203" score="0">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="165" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="101" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="48.0916500000000013"/>
            </Node>
           </Node>
           <Node id="69" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="45">
           <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
           <Node id="70">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="102">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="7.98750000000000071"/>
             <Node id="126">
              <SimplePredicate field="Age" operator="lessOrEqual" value="41.5"/>
              <Node id="146" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
              </Node>
              <Node id="147" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
              </Node>
             </Node>
             <Node id="127" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="41.5"/>
             </Node>
            </Node>
            <Node id="103" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="7.98750000000000071"/>
            </Node>
           </Node>
           <Node id="71" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="3">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
        <Array n="2" type="string">&quot;&quot;   &quot;C&quot;</Array>
       </SimpleSetPredicate>
       <Node id="4">
        <SimplePredicate field="Fare" operator="lessOrEqual" value="9.97710000000000008"/>
        <Node id="8">
         <SimplePredicate field="Age" operator="lessOrEqual" value="21"/>
         <Node id="16" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
         </Node>
         <Node id="17">
          <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
          <Node id="30" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="5.6208499999999999"/>
          </Node>
          <Node id="31" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="5.6208499999999999"/>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Age" operator="greaterThan" value="21"/>
         <Node id="18">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="7.5625"/>
          <Node id="32" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="6.83124999999999982"/>
          </Node>
          <Node id="33">
           <SimplePredicate field="Fare" operator="greaterThan" value="6.83124999999999982"/>
           <Node id="52" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="53">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
            <Node id="68" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="22.75"/>
            </Node>
            <Node id="69" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="22.75"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Fare" operator="greaterThan" value="7.5625"/>
          <Node id="34" score="0">
           <SimplePredicate field="Age" operator="lessOrEqual" value="31"/>
          </Node>
          <Node id="35" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="31"/>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Fare" operator="greaterThan" value="9.97710000000000008"/>
        <Node id="10">
         <SimplePredicate field="Age" operator="lessOrEqual" value="66.5"/>
         <Node id="20">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="36">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="54">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="70" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="47"/>
            </Node>
            <Node id="71" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="47"/>
            </Node>
           </Node>
           <Node id="55">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            <Node id="72" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2"/>
            </Node>
            <Node id="73">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2"/>
             <Node id="88" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="89" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="37" score="1">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
          </Node>
         </Node>
         <Node id="21" score="0">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
         </Node>
        </Node>
        <Node id="11" score="0">
         <SimplePredicate field="Age" operator="greaterThan" value="66.5"/>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
        <Array n="2" type="string">&quot;Q&quot;   &quot;S&quot;</Array>
       </SimpleSetPredicate>
       <Node id="6">
        <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
        <Node id="12">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
          <Node id="38" score="0">
           <SimplePredicate field="Age" operator="lessOrEqual" value="8.5"/>
          </Node>
          <Node id="39" score="1">
           <SimplePredicate field="Age" operator="greaterThan" value="8.5"/>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          <Node id="40" score="1">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="13.75"/>
          </Node>
          <Node id="41" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="13.75"/>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="24">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="42">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="26.1437500000000007"/>
           <Node id="56" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="57" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="43">
           <SimplePredicate field="Fare" operator="greaterThan" value="26.1437500000000007"/>
           <Node id="58">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="30.25"/>
            <Node id="74" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="75" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="59">
            <SimplePredicate field="Fare" operator="greaterThan" value="30.25"/>
            <Node id="76">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
             <Node id="90" score="1">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="91" score="0">
              <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
            <Node id="77">
             <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
             <Node id="92" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="37.5"/>
             </Node>
             <Node id="93">
              <SimplePredicate field="Age" operator="greaterThan" value="37.5"/>
              <Node id="100" score="0">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="101">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="108" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="46"/>
               </Node>
               <Node id="109" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="46"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="25">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="44" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="15"/>
          </Node>
          <Node id="45" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="15"/>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
        <Node id="14">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="7.73124999999999929"/>
         <Node id="26">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="46" score="0">
           <SimplePredicate field="Age" operator="lessOrEqual" value="20"/>
          </Node>
          <Node id="47" score="1">
           <SimplePredicate field="Age" operator="greaterThan" value="20"/>
          </Node>
         </Node>
         <Node id="27" score="0">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
         </Node>
        </Node>
        <Node id="15">
         <SimplePredicate field="Fare" operator="greaterThan" value="7.73124999999999929"/>
         <Node id="28">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="10.8249999999999993"/>
          <Node id="48">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.88750000000000018"/>
           <Node id="60">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="78" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="79" score="0">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="49">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.88750000000000018"/>
           <Node id="62">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
            <Node id="80" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="81" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
           <Node id="63">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
            <Node id="82">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1"/>
             <Node id="94" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.9104000000000001"/>
             </Node>
             <Node id="95" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.9104000000000001"/>
             </Node>
            </Node>
            <Node id="83" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="29">
          <SimplePredicate field="Fare" operator="greaterThan" value="10.8249999999999993"/>
          <Node id="50">
           <SimplePredicate field="Age" operator="lessOrEqual" value="6.5"/>
           <Node id="64" score="0">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;Q&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="65">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="84" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
            </Node>
            <Node id="85">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
             <Node id="96">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
              <Node id="102">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;female&quot;</Array>
               </SimpleSetPredicate>
               <Node id="110" score="0">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
               </Node>
               <Node id="111">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                <Node id="122">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
                 <Node id="130" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
                 </Node>
                 <Node id="131" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
                 </Node>
                </Node>
                <Node id="123" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="103">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;male&quot;</Array>
               </SimpleSetPredicate>
               <Node id="112" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
               </Node>
               <Node id="113" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="97" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="51">
           <SimplePredicate field="Age" operator="greaterThan" value="6.5"/>
           <Node id="66">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
            <Node id="86">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="51.6979000000000042"/>
             <Node id="98">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
              <Node id="104">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="26.6374999999999993"/>
               <Node id="114">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                <Node id="124" score="1">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="37.5"/>
                </Node>
                <Node id="125" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="37.5"/>
                </Node>
               </Node>
               <Node id="115">
                <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                <Node id="126" score="0">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;Q&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
                <Node id="127" score="0">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
               </Node>
              </Node>
              <Node id="105">
               <SimplePredicate field="Fare" operator="greaterThan" value="26.6374999999999993"/>
               <Node id="116" score="0">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="4"/>
               </Node>
               <Node id="117">
                <SimplePredicate field="Parch" operator="greaterThan" value="4"/>
                <Node id="128" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
                </Node>
                <Node id="129">
                 <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
                 <Node id="132" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="40.5"/>
                 </Node>
                 <Node id="133" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="40.5"/>
                 </Node>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="99">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
              <Node id="106">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
               <Node id="118" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;Q&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="119" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
              <Node id="107">
               <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
               <Node id="120" score="1">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;Q&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="121" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="87" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="51.6979000000000042"/>
            </Node>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="4">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
       <Node id="4">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;female&quot;</Array>
        </SimpleSetPredicate>
        <Node id="8">
         <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
         <Node id="16">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="29.3916499999999985"/>
          <Node id="28">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="48">
            <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
            <Node id="74" score="1">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;Q&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="75" score="1">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
           <Node id="49">
            <SimplePredicate field="Age" operator="greaterThan" value="37"/>
            <Node id="76" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
            </Node>
            <Node id="77">
             <SimplePredicate field="Age" operator="greaterThan" value="39"/>
             <Node id="92" score="1">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="93">
              <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
              <Node id="106" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="12"/>
              </Node>
              <Node id="107" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="12"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="29" score="1">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
          </Node>
         </Node>
         <Node id="17" score="1">
          <SimplePredicate field="Fare" operator="greaterThan" value="29.3916499999999985"/>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
         <Node id="18" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
         </Node>
         <Node id="19">
          <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
          <Node id="30" score="1">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           <Node id="50" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="135.775000000000006"/>
           </Node>
           <Node id="51" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="135.775000000000006"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;male&quot;</Array>
        </SimpleSetPredicate>
        <Node id="10">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="1" type="string">&quot;C&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20">
          <SimplePredicate field="Age" operator="lessOrEqual" value="50"/>
          <Node id="32">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="28.7250000000000014"/>
           <Node id="52">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="27.1354000000000006"/>
            <Node id="78" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="25.2749999999999986"/>
            </Node>
            <Node id="79" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="25.2749999999999986"/>
            </Node>
           </Node>
           <Node id="53" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="27.1354000000000006"/>
           </Node>
          </Node>
          <Node id="33">
           <SimplePredicate field="Fare" operator="greaterThan" value="28.7250000000000014"/>
           <Node id="54">
            <SimplePredicate field="Age" operator="lessOrEqual" value="38"/>
            <Node id="80">
             <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
             <Node id="94">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="123.258299999999991"/>
              <Node id="108" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
              </Node>
              <Node id="109" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
              </Node>
             </Node>
             <Node id="95" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="123.258299999999991"/>
             </Node>
            </Node>
            <Node id="81">
             <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
             <Node id="96" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
             </Node>
             <Node id="97">
              <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
              <Node id="110">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="30.5978999999999992"/>
               <Node id="120" score="1">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               </Node>
               <Node id="121" score="0">
                <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
              <Node id="111" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="30.5978999999999992"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="55" score="1">
            <SimplePredicate field="Age" operator="greaterThan" value="38"/>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Age" operator="greaterThan" value="50"/>
          <Node id="34" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="70.2896000000000072"/>
          </Node>
          <Node id="35">
           <SimplePredicate field="Fare" operator="greaterThan" value="70.2896000000000072"/>
           <Node id="56" score="1">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="57" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="14.25"/>
          <Node id="36">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="12.7624999999999993"/>
           <Node id="58" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="7.75"/>
           </Node>
           <Node id="59" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="7.75"/>
           </Node>
          </Node>
          <Node id="37">
           <SimplePredicate field="Fare" operator="greaterThan" value="12.7624999999999993"/>
           <Node id="60" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Fare" operator="greaterThan" value="14.25"/>
          <Node id="38">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
           <Node id="62" score="0">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;Q&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="63">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="82">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             <Node id="98">
              <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
              <Node id="112">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="30.75"/>
               <Node id="122" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="46"/>
               </Node>
               <Node id="123" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="46"/>
               </Node>
              </Node>
              <Node id="113">
               <SimplePredicate field="Fare" operator="greaterThan" value="30.75"/>
               <Node id="124" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="35.25"/>
               </Node>
               <Node id="125">
                <SimplePredicate field="Fare" operator="greaterThan" value="35.25"/>
                <Node id="130">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="48.5"/>
                 <Node id="132">
                  <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                  <Node id="134" score="0">
                   <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                  </Node>
                  <Node id="135" score="0">
                   <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                  </Node>
                 </Node>
                 <Node id="133">
                  <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                  <Node id="136" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="23"/>
                  </Node>
                  <Node id="137">
                   <SimplePredicate field="Age" operator="greaterThan" value="23"/>
                   <Node id="138">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="39.5"/>
                    <Node id="140">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
                     <Node id="142" score="1">
                      <SimplePredicate field="Age" operator="lessOrEqual" value="28"/>
                     </Node>
                     <Node id="143">
                      <SimplePredicate field="Age" operator="greaterThan" value="28"/>
                      <Node id="144" score="0">
                       <SimplePredicate field="Fare" operator="lessOrEqual" value="99.4249999999999972"/>
                      </Node>
                      <Node id="145" score="1">
                       <SimplePredicate field="Fare" operator="greaterThan" value="99.4249999999999972"/>
                      </Node>
                     </Node>
                    </Node>
                    <Node id="141" score="1">
                     <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
                    </Node>
                   </Node>
                   <Node id="139" score="0">
                    <SimplePredicate field="Age" operator="greaterThan" value="39.5"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="131" score="1">
                 <SimplePredicate field="Age" operator="greaterThan" value="48.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="99" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="53"/>
             </Node>
            </Node>
            <Node id="83">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
             <Node id="100" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="101">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              <Node id="114" score="1">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="115" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="39" score="0">
           <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
       <Node id="6">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
        </SimpleSetPredicate>
        <Node id="12">
         <SimplePredicate field="Age" operator="lessOrEqual" value="16.5"/>
         <Node id="24">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="40">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;C&quot;</Array>
           </SimpleSetPredicate>
           <Node id="64" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8479500000000009"/>
           </Node>
           <Node id="65">
            <SimplePredicate field="Fare" operator="greaterThan" value="12.8479500000000009"/>
            <Node id="84" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="16.6208500000000008"/>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="16.6208500000000008"/>
            </Node>
           </Node>
          </Node>
          <Node id="41" score="1">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
          </Node>
         </Node>
         <Node id="25">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="42">
           <SimplePredicate field="Age" operator="lessOrEqual" value="6"/>
           <Node id="66" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="3"/>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="3"/>
           </Node>
          </Node>
          <Node id="43" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="6"/>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimplePredicate field="Age" operator="greaterThan" value="16.5"/>
         <Node id="26">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="44">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
           <Node id="68">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="14.8520499999999984"/>
            <Node id="86">
             <SimplePredicate field="Age" operator="lessOrEqual" value="28.25"/>
             <Node id="102">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
              <Node id="116">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="126" score="0">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="127" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
              <Node id="117" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="103">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
              <Node id="118">
               <SimplePredicate field="Age" operator="lessOrEqual" value="22.75"/>
               <Node id="128" score="0">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="129" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
              <Node id="119" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="22.75"/>
              </Node>
             </Node>
            </Node>
            <Node id="87" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="28.25"/>
            </Node>
           </Node>
           <Node id="69" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="14.8520499999999984"/>
           </Node>
          </Node>
          <Node id="45" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
          </Node>
         </Node>
         <Node id="27">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="46">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="70" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="7.24164999999999992"/>
           </Node>
           <Node id="71">
            <SimplePredicate field="Fare" operator="greaterThan" value="7.24164999999999992"/>
            <Node id="88" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
            </Node>
            <Node id="89" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="30"/>
            </Node>
           </Node>
          </Node>
          <Node id="47">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
           <Node id="72">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="19.375"/>
            <Node id="90">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="104" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
             </Node>
             <Node id="105" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
             </Node>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="73" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="19.375"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
        </SimpleSetPredicate>
        <Node id="14" score="0">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
        </Node>
        <Node id="15" score="0">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="5">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Fare" operator="lessOrEqual" value="74.375"/>
       <Node id="4">
        <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
        <Node id="8">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="13.6458499999999994"/>
         <Node id="14">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="24">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8249999999999993"/>
           <Node id="42" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="11.25"/>
           </Node>
           <Node id="43" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="11.25"/>
           </Node>
          </Node>
          <Node id="25" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="12.8249999999999993"/>
          </Node>
         </Node>
         <Node id="15">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
          <Node id="26">
           <SimplePredicate field="Age" operator="lessOrEqual" value="40.5"/>
           <Node id="44">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="72" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="20"/>
            </Node>
            <Node id="73" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="20"/>
            </Node>
           </Node>
           <Node id="45" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="27">
           <SimplePredicate field="Age" operator="greaterThan" value="40.5"/>
           <Node id="46" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="45"/>
           </Node>
           <Node id="47" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="45"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Fare" operator="greaterThan" value="13.6458499999999994"/>
         <Node id="16">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="28">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="45.5396000000000001"/>
           <Node id="48">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="74" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="75">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="104" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="48"/>
             </Node>
             <Node id="105" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="48"/>
             </Node>
            </Node>
           </Node>
           <Node id="49">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="76" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
            </Node>
            <Node id="77" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="39"/>
            </Node>
           </Node>
          </Node>
          <Node id="29">
           <SimplePredicate field="Fare" operator="greaterThan" value="45.5396000000000001"/>
           <Node id="50" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="51">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
            <Node id="78" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="79" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="17">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="30" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="14.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Age" operator="greaterThan" value="14.5"/>
           <Node id="52">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="80">
             <SimplePredicate field="Age" operator="lessOrEqual" value="43"/>
             <Node id="106">
              <SimplePredicate field="Age" operator="lessOrEqual" value="39.5"/>
              <Node id="134">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="158" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="159">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
                <Node id="186">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 <Node id="202" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="42.75"/>
                 </Node>
                 <Node id="203" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="42.75"/>
                 </Node>
                </Node>
                <Node id="187">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 <Node id="204">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="34"/>
                  <Node id="214" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="55.0499999999999972"/>
                  </Node>
                  <Node id="215" score="1">
                   <SimplePredicate field="Fare" operator="greaterThan" value="55.0499999999999972"/>
                  </Node>
                 </Node>
                 <Node id="205" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="34"/>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="135">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               <Node id="160">
                <SimplePredicate field="Age" operator="lessOrEqual" value="31"/>
                <Node id="188" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                </Node>
                <Node id="189">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 <Node id="206">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="49.75"/>
                  <Node id="216" score="1">
                   <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;female&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="217" score="0">
                   <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;male&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                 </Node>
                 <Node id="207" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="49.75"/>
                 </Node>
                </Node>
               </Node>
               <Node id="161" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="31"/>
               </Node>
              </Node>
             </Node>
             <Node id="107" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="39.5"/>
             </Node>
            </Node>
            <Node id="81">
             <SimplePredicate field="Age" operator="greaterThan" value="43"/>
             <Node id="108">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              <Node id="136">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="162">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                <Node id="190">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="31.4103999999999992"/>
                 <Node id="208" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="28.2749999999999986"/>
                 </Node>
                 <Node id="209" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="28.2749999999999986"/>
                 </Node>
                </Node>
                <Node id="191" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="31.4103999999999992"/>
                </Node>
               </Node>
               <Node id="163" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
              <Node id="137">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               <Node id="164" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="54.5"/>
               </Node>
               <Node id="165" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="54.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="109" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="53">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            <Node id="82">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             <Node id="110" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="58.5"/>
             </Node>
             <Node id="111" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="58.5"/>
             </Node>
            </Node>
            <Node id="83">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
             <Node id="112" score="1">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="113" score="0">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
        <Node id="10">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
         </SimpleSetPredicate>
         <Node id="18">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
          <Node id="32">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="54">
            <SimplePredicate field="Age" operator="lessOrEqual" value="29.25"/>
            <Node id="84">
             <SimplePredicate field="Age" operator="lessOrEqual" value="14.75"/>
             <Node id="114" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="13.9354499999999994"/>
             </Node>
             <Node id="115">
              <SimplePredicate field="Fare" operator="greaterThan" value="13.9354499999999994"/>
              <Node id="138" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="7"/>
              </Node>
              <Node id="139" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="7"/>
              </Node>
             </Node>
            </Node>
            <Node id="85">
             <SimplePredicate field="Age" operator="greaterThan" value="14.75"/>
             <Node id="116">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
              <Node id="140" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="16"/>
              </Node>
              <Node id="141">
               <SimplePredicate field="Age" operator="greaterThan" value="16"/>
               <Node id="166">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                <Node id="192" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="10.84375"/>
                </Node>
                <Node id="193" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="10.84375"/>
                </Node>
               </Node>
               <Node id="167" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="117" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="29.25"/>
           </Node>
          </Node>
          <Node id="33">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
           <Node id="56">
            <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
            <Node id="86">
             <SimplePredicate field="Age" operator="lessOrEqual" value="28.75"/>
             <Node id="118">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
              <Node id="142">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="168">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
                <Node id="194">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="16.6229000000000013"/>
                 <Node id="210" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="16"/>
                 </Node>
                 <Node id="211" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="16"/>
                 </Node>
                </Node>
                <Node id="195" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="16.6229000000000013"/>
                </Node>
               </Node>
               <Node id="169">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="196" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
                </Node>
                <Node id="197">
                 <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
                 <Node id="212" score="0">
                  <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 </Node>
                 <Node id="213" score="0">
                  <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="143">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
               <Node id="170" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="11.2374999999999989"/>
               </Node>
               <Node id="171" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="11.2374999999999989"/>
               </Node>
              </Node>
             </Node>
             <Node id="119" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
             </Node>
            </Node>
            <Node id="87" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="28.75"/>
            </Node>
           </Node>
           <Node id="57" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="19" score="1">
          <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="17.25"/>
          <Node id="34">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
           <Node id="58" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="59">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
            <Node id="88" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="8.13540000000000063"/>
            </Node>
            <Node id="89">
             <SimplePredicate field="Fare" operator="greaterThan" value="8.13540000000000063"/>
             <Node id="120" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="121" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="35">
           <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
           <Node id="60">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="90">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
             <Node id="122">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="10.7979000000000003"/>
              <Node id="144">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="9.90625"/>
               <Node id="172" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="173" score="0">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
              <Node id="145" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="9.90625"/>
              </Node>
             </Node>
             <Node id="123" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="10.7979000000000003"/>
             </Node>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
            </Node>
           </Node>
           <Node id="61">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="92" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="33.75"/>
            </Node>
            <Node id="93" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="33.75"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Fare" operator="greaterThan" value="17.25"/>
          <Node id="36">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="4.5"/>
           <Node id="62">
            <SimplePredicate field="Age" operator="lessOrEqual" value="6"/>
            <Node id="94">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
             <Node id="124" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="125">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
              <Node id="146" score="1">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="147">
               <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
               <Node id="174" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
               </Node>
               <Node id="175" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="95" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
            </Node>
           </Node>
           <Node id="63">
            <SimplePredicate field="Age" operator="greaterThan" value="6"/>
            <Node id="96" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
            </Node>
            <Node id="97">
             <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
             <Node id="126">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
              <Node id="148">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="176" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
               </Node>
               <Node id="177">
                <SimplePredicate field="Age" operator="greaterThan" value="30"/>
                <Node id="198" score="1">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
                </Node>
                <Node id="199" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="149">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="178">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
                <Node id="200" score="1">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="40.5"/>
                </Node>
                <Node id="201" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="40.5"/>
                </Node>
               </Node>
               <Node id="179" score="0">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
             </Node>
             <Node id="127" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="4.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Fare" operator="greaterThan" value="74.375"/>
       <Node id="6" score="0">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="1" type="string">&quot;Q&quot;</Array>
        </SimpleSetPredicate>
       </Node>
       <Node id="7">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
        </SimpleSetPredicate>
        <Node id="12">
         <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
         <Node id="22">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="38">
           <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
           <Node id="64">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="98" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="99" score="0">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
           <Node id="65" score="1">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
           <Node id="66" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="78.7333500000000015"/>
           </Node>
           <Node id="67">
            <SimplePredicate field="Fare" operator="greaterThan" value="78.7333500000000015"/>
            <Node id="100">
             <SimplePredicate field="Age" operator="lessOrEqual" value="44"/>
             <Node id="128">
              <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
              <Node id="150">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="180" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="181" score="0">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
              <Node id="151" score="1">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="129" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
             </Node>
            </Node>
            <Node id="101">
             <SimplePredicate field="Age" operator="greaterThan" value="44"/>
             <Node id="130" score="1">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="131">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
              <Node id="152" score="0">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="153">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="182" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="99.9937500000000057"/>
               </Node>
               <Node id="183" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="99.9937500000000057"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="23">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="40">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="79.25"/>
           <Node id="68" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="69" score="0">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
           </Node>
          </Node>
          <Node id="41">
           <SimplePredicate field="Fare" operator="greaterThan" value="79.25"/>
           <Node id="70" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="143.591650000000016"/>
           </Node>
           <Node id="71">
            <SimplePredicate field="Fare" operator="greaterThan" value="143.591650000000016"/>
            <Node id="102">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="159.164600000000007"/>
             <Node id="132">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="152.506250000000023"/>
              <Node id="154">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;female&quot;</Array>
               </SimpleSetPredicate>
               <Node id="184" score="1">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
               </Node>
               <Node id="185" score="0">
                <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
               </Node>
              </Node>
              <Node id="155" score="1">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;male&quot;</Array>
               </SimpleSetPredicate>
              </Node>
             </Node>
             <Node id="133">
              <SimplePredicate field="Fare" operator="greaterThan" value="152.506250000000023"/>
              <Node id="156" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
              </Node>
              <Node id="157" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="39"/>
              </Node>
             </Node>
            </Node>
            <Node id="103" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="159.164600000000007"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="13" score="0">
         <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="6">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Age" operator="lessOrEqual" value="6.5"/>
       <Node id="4">
        <SimplePredicate field="Fare" operator="lessOrEqual" value="20.1666499999999971"/>
        <Node id="8" score="1">
         <SimplePredicate field="Age" operator="lessOrEqual" value="1.5"/>
        </Node>
        <Node id="9">
         <SimplePredicate field="Age" operator="greaterThan" value="1.5"/>
         <Node id="16" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
         </Node>
         <Node id="17" score="1">
          <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Fare" operator="greaterThan" value="20.1666499999999971"/>
        <Node id="10" score="1">
         <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
        </Node>
        <Node id="11">
         <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
         <Node id="18" score="0">
          <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
         </Node>
         <Node id="19" score="1">
          <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Age" operator="greaterThan" value="6.5"/>
       <Node id="6">
        <SimplePredicate field="Fare" operator="lessOrEqual" value="50.9874999999999972"/>
        <Node id="12">
         <SimplePredicate field="Age" operator="lessOrEqual" value="75.5"/>
         <Node id="20">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="26">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;Q&quot;</Array>
           </SimpleSetPredicate>
           <Node id="34" score="1">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="35">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            <Node id="46">
             <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
             <Node id="62" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.68125000000000036"/>
             </Node>
             <Node id="63" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.68125000000000036"/>
             </Node>
            </Node>
            <Node id="47" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="30"/>
            </Node>
           </Node>
          </Node>
          <Node id="27">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="36">
            <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
            <Node id="48">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
             <Node id="64">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="26.125"/>
              <Node id="80">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="100" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8249999999999993"/>
               </Node>
               <Node id="101">
                <SimplePredicate field="Fare" operator="greaterThan" value="12.8249999999999993"/>
                <Node id="128">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="19.5"/>
                 <Node id="150" score="0">
                  <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
                 </Node>
                 <Node id="151" score="1">
                  <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
                 </Node>
                </Node>
                <Node id="129" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="19.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="81">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="102" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="19.875"/>
               </Node>
               <Node id="103" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="19.875"/>
               </Node>
              </Node>
             </Node>
             <Node id="65" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="26.125"/>
             </Node>
            </Node>
            <Node id="49">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             <Node id="66">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="82" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.40000000000000036"/>
              </Node>
              <Node id="83">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.40000000000000036"/>
               <Node id="104">
                <SimplePredicate field="Age" operator="lessOrEqual" value="19.5"/>
                <Node id="130">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
                 <Node id="152" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8479500000000009"/>
                 </Node>
                 <Node id="153">
                  <SimplePredicate field="Fare" operator="greaterThan" value="12.8479500000000009"/>
                  <Node id="164" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="14.75"/>
                  </Node>
                  <Node id="165">
                   <SimplePredicate field="Age" operator="greaterThan" value="14.75"/>
                   <Node id="170">
                    <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                     <Array n="1" type="string">&quot;C&quot;</Array>
                    </SimpleSetPredicate>
                    <Node id="174" score="1">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="16"/>
                    </Node>
                    <Node id="175" score="0">
                     <SimplePredicate field="Age" operator="greaterThan" value="16"/>
                    </Node>
                   </Node>
                   <Node id="171" score="0">
                    <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                     <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                    </SimpleSetPredicate>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="131" score="1">
                 <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
                </Node>
               </Node>
               <Node id="105" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="19.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="67">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              <Node id="84">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="13.4375"/>
               <Node id="106" score="1">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="107">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                <Node id="132" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="9.19374999999999964"/>
                </Node>
                <Node id="133" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="9.19374999999999964"/>
                </Node>
               </Node>
              </Node>
              <Node id="85">
               <SimplePredicate field="Fare" operator="greaterThan" value="13.4375"/>
               <Node id="108">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="20.1666499999999971"/>
                <Node id="134" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="15.9728999999999992"/>
                </Node>
                <Node id="135" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="15.9728999999999992"/>
                </Node>
               </Node>
               <Node id="109" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="20.1666499999999971"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="37">
            <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
            <Node id="50">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
             <Node id="68" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="11.75"/>
             </Node>
             <Node id="69">
              <SimplePredicate field="Fare" operator="greaterThan" value="11.75"/>
              <Node id="86">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="110" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="49"/>
               </Node>
               <Node id="111">
                <SimplePredicate field="Age" operator="greaterThan" value="49"/>
                <Node id="136" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="27.6312500000000014"/>
                </Node>
                <Node id="137" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="27.6312500000000014"/>
                </Node>
               </Node>
              </Node>
              <Node id="87">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               <Node id="112" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="43"/>
               </Node>
               <Node id="113" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="43"/>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="51">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             <Node id="70">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="32.8812500000000014"/>
              <Node id="88" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
              </Node>
              <Node id="89" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="39"/>
              </Node>
             </Node>
             <Node id="71" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="32.8812500000000014"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
          <Node id="28">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
           <Node id="38">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="26.3187500000000014"/>
            <Node id="52" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="7.13335000000000008"/>
            </Node>
            <Node id="53">
             <SimplePredicate field="Fare" operator="greaterThan" value="7.13335000000000008"/>
             <Node id="72">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.23959999999999937"/>
              <Node id="90">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="114" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="115" score="1">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
              <Node id="91" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="73">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.23959999999999937"/>
              <Node id="92">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="116" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
               </Node>
               <Node id="117">
                <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                <Node id="138" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
                </Node>
                <Node id="139" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
                </Node>
               </Node>
              </Node>
              <Node id="93">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="118">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
                </SimpleSetPredicate>
                <Node id="140">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="22.4646000000000008"/>
                 <Node id="154" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8479500000000009"/>
                 </Node>
                 <Node id="155">
                  <SimplePredicate field="Fare" operator="greaterThan" value="12.8479500000000009"/>
                  <Node id="166" score="1">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="23.5"/>
                  </Node>
                  <Node id="167">
                   <SimplePredicate field="Age" operator="greaterThan" value="23.5"/>
                   <Node id="172">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="15.3728999999999996"/>
                    <Node id="176" score="0">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
                    </Node>
                    <Node id="177" score="1">
                     <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
                    </Node>
                   </Node>
                   <Node id="173" score="0">
                    <SimplePredicate field="Fare" operator="greaterThan" value="15.3728999999999996"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="141">
                 <SimplePredicate field="Fare" operator="greaterThan" value="22.4646000000000008"/>
                 <Node id="156" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="29"/>
                 </Node>
                 <Node id="157" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="29"/>
                 </Node>
                </Node>
               </Node>
               <Node id="119">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="142">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
                 <Node id="158" score="0">
                  <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
                 </Node>
                 <Node id="159" score="0">
                  <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                 </Node>
                </Node>
                <Node id="143" score="0">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="39">
            <SimplePredicate field="Fare" operator="greaterThan" value="26.3187500000000014"/>
            <Node id="54">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             <Node id="74">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="27.1354000000000006"/>
              <Node id="94">
               <SimplePredicate field="Age" operator="lessOrEqual" value="53.5"/>
               <Node id="120" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="40.5"/>
               </Node>
               <Node id="121" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="40.5"/>
               </Node>
              </Node>
              <Node id="95" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="53.5"/>
              </Node>
             </Node>
             <Node id="75">
              <SimplePredicate field="Fare" operator="greaterThan" value="27.1354000000000006"/>
              <Node id="96" score="0">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="97">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="122" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
               </Node>
               <Node id="123">
                <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
                <Node id="144">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                 <Node id="160" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="30.75"/>
                 </Node>
                 <Node id="161" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="30.75"/>
                 </Node>
                </Node>
                <Node id="145" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="55" score="0">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="29" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
        </Node>
        <Node id="13" score="1">
         <SimplePredicate field="Age" operator="greaterThan" value="75.5"/>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Fare" operator="greaterThan" value="50.9874999999999972"/>
        <Node id="14">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22" score="1">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
         </Node>
         <Node id="23">
          <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
          <Node id="30" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
           <Node id="40" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="67.2750000000000057"/>
           </Node>
           <Node id="41">
            <SimplePredicate field="Fare" operator="greaterThan" value="67.2750000000000057"/>
            <Node id="56" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
            </Node>
            <Node id="57" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="32"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="24">
          <SimplePredicate field="Age" operator="lessOrEqual" value="50.5"/>
          <Node id="32">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;C&quot;</Array>
           </SimpleSetPredicate>
           <Node id="42" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="77.9646000000000043"/>
           </Node>
           <Node id="43">
            <SimplePredicate field="Fare" operator="greaterThan" value="77.9646000000000043"/>
            <Node id="58">
             <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
             <Node id="76" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="77" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="59" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="32"/>
            </Node>
           </Node>
          </Node>
          <Node id="33">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="44">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="143.556250000000006"/>
            <Node id="60">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="5"/>
             <Node id="78">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              <Node id="98">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
               <Node id="124">
                <SimplePredicate field="Age" operator="lessOrEqual" value="37.5"/>
                <Node id="146">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 <Node id="162">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
                  <Node id="168" score="0">
                   <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                  </Node>
                  <Node id="169" score="0">
                   <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                  </Node>
                 </Node>
                 <Node id="163" score="1">
                  <SimplePredicate field="Age" operator="greaterThan" value="30"/>
                 </Node>
                </Node>
                <Node id="147" score="0">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
               <Node id="125" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="37.5"/>
               </Node>
              </Node>
              <Node id="99">
               <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
               <Node id="126">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                <Node id="148" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="111.825000000000003"/>
                </Node>
                <Node id="149" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="111.825000000000003"/>
                </Node>
               </Node>
               <Node id="127" score="0">
                <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="79" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
            <Node id="61" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="5"/>
            </Node>
           </Node>
           <Node id="45" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="143.556250000000006"/>
           </Node>
          </Node>
         </Node>
         <Node id="25" score="0">
          <SimplePredicate field="Age" operator="greaterThan" value="50.5"/>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="7">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;female&quot;</Array>
       </SimpleSetPredicate>
       <Node id="4">
        <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
        <Node id="8">
         <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
         <Node id="16">
          <SimplePredicate field="Age" operator="lessOrEqual" value="43.5"/>
          <Node id="30" score="1">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           <Node id="52">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="22"/>
            <Node id="80">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
             <Node id="108" score="1">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="109" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="81">
             <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
             <Node id="110" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;Q&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="111">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
              <Node id="122" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
              </Node>
              <Node id="123">
               <SimplePredicate field="Age" operator="greaterThan" value="37"/>
               <Node id="128" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
               </Node>
               <Node id="129" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="39"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="53" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="22"/>
           </Node>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Age" operator="greaterThan" value="43.5"/>
          <Node id="32">
           <SimplePredicate field="Age" operator="lessOrEqual" value="44.5"/>
           <Node id="54" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="33">
           <SimplePredicate field="Age" operator="greaterThan" value="44.5"/>
           <Node id="56" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="56.5"/>
           </Node>
           <Node id="57">
            <SimplePredicate field="Age" operator="greaterThan" value="56.5"/>
            <Node id="82" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="83" score="0">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
         <Node id="18" score="0">
          <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
         </Node>
         <Node id="19">
          <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
          <Node id="34" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
          </Node>
          <Node id="35">
           <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
           <Node id="58" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="59">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="84">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
             <Node id="112" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="135.775000000000006"/>
             </Node>
             <Node id="113" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="135.775000000000006"/>
             </Node>
            </Node>
            <Node id="85" score="1">
             <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
        <Node id="10">
         <SimplePredicate field="Age" operator="lessOrEqual" value="38.5"/>
         <Node id="20">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="24.8083499999999972"/>
          <Node id="36">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
           <Node id="60" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="7"/>
           </Node>
           <Node id="61">
            <SimplePredicate field="Age" operator="greaterThan" value="7"/>
            <Node id="86" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="14.75"/>
            </Node>
            <Node id="87">
             <SimplePredicate field="Age" operator="greaterThan" value="14.75"/>
             <Node id="114">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.88750000000000018"/>
              <Node id="124">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="130" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="6.98749999999999982"/>
               </Node>
               <Node id="131">
                <SimplePredicate field="Fare" operator="greaterThan" value="6.98749999999999982"/>
                <Node id="134" score="1">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
                </Node>
                <Node id="135" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="125" score="1">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="115">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.88750000000000018"/>
              <Node id="126">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
               <Node id="132">
                <SimplePredicate field="Age" operator="lessOrEqual" value="23.5"/>
                <Node id="136" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="9.83334999999999937"/>
                </Node>
                <Node id="137">
                 <SimplePredicate field="Fare" operator="greaterThan" value="9.83334999999999937"/>
                 <Node id="138" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="10.1791999999999998"/>
                 </Node>
                 <Node id="139">
                  <SimplePredicate field="Fare" operator="greaterThan" value="10.1791999999999998"/>
                  <Node id="140" score="0">
                   <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                  </Node>
                  <Node id="141">
                   <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                   <Node id="142" score="1">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="16.1270999999999987"/>
                   </Node>
                   <Node id="143" score="0">
                    <SimplePredicate field="Fare" operator="greaterThan" value="16.1270999999999987"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="133" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="23.5"/>
               </Node>
              </Node>
              <Node id="127" score="1">
               <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Fare" operator="greaterThan" value="24.8083499999999972"/>
          <Node id="38">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
           <Node id="62">
            <SimplePredicate field="Age" operator="lessOrEqual" value="7"/>
            <Node id="88" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
            </Node>
            <Node id="89" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
            </Node>
           </Node>
           <Node id="63" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="7"/>
           </Node>
          </Node>
          <Node id="39" score="1">
           <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimplePredicate field="Age" operator="greaterThan" value="38.5"/>
         <Node id="22">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
          <Node id="40" score="0">
           <SimplePredicate field="Age" operator="lessOrEqual" value="54"/>
          </Node>
          <Node id="41" score="1">
           <SimplePredicate field="Age" operator="greaterThan" value="54"/>
          </Node>
         </Node>
         <Node id="23" score="0">
          <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;male&quot;</Array>
       </SimpleSetPredicate>
       <Node id="6">
        <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
        <Node id="12">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="26.1437500000000007"/>
         <Node id="24">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="42" score="0">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
          </Node>
          <Node id="43" score="0">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
         <Node id="25">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="44">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="11"/>
           <Node id="64" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="65">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="90" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="7.76250000000000018"/>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="7.76250000000000018"/>
            </Node>
           </Node>
          </Node>
          <Node id="45">
           <SimplePredicate field="Fare" operator="greaterThan" value="11"/>
           <Node id="66" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimplePredicate field="Fare" operator="greaterThan" value="26.1437500000000007"/>
         <Node id="26" score="1">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="26.4187500000000028"/>
         </Node>
         <Node id="27">
          <SimplePredicate field="Fare" operator="greaterThan" value="26.4187500000000028"/>
          <Node id="46">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="56.1979000000000042"/>
           <Node id="68" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="69" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="47">
           <SimplePredicate field="Fare" operator="greaterThan" value="56.1979000000000042"/>
           <Node id="70">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="92" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="64.9979000000000013"/>
            </Node>
            <Node id="93" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="64.9979000000000013"/>
            </Node>
           </Node>
           <Node id="71">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="94" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="95">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             <Node id="116" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;Q&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="117" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
        <Node id="14">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="387.664600000000007"/>
         <Node id="28">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
          <Node id="48">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="72">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="70.3229000000000042"/>
            <Node id="96" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="35.0020999999999987"/>
            </Node>
            <Node id="97" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="35.0020999999999987"/>
            </Node>
           </Node>
           <Node id="73">
            <SimplePredicate field="Fare" operator="greaterThan" value="70.3229000000000042"/>
            <Node id="98" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="12.5"/>
            </Node>
            <Node id="99" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="12.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="49">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="74">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="100" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="60.3896000000000015"/>
            </Node>
            <Node id="101" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="60.3896000000000015"/>
            </Node>
           </Node>
           <Node id="75">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="102" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="15"/>
            </Node>
            <Node id="103" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="15"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="29">
          <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          <Node id="50">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="20.8249999999999993"/>
           <Node id="76">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="104" score="0">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="105">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="118" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="12.5"/>
             </Node>
             <Node id="119" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="12.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="77" score="1">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="51">
           <SimplePredicate field="Fare" operator="greaterThan" value="20.8249999999999993"/>
           <Node id="78">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
            <Node id="106" score="0">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            </Node>
            <Node id="107">
             <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             <Node id="120" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
             </Node>
             <Node id="121" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="79" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15" score="1">
         <SimplePredicate field="Fare" operator="greaterThan" value="387.664600000000007"/>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="8">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;female&quot;</Array>
       </SimpleSetPredicate>
       <Node id="4">
        <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
        <Node id="8">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="25.0749999999999993"/>
         <Node id="16">
          <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
          <Node id="28">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="48">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="70">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             <Node id="96">
              <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
              <Node id="126">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
               <Node id="152" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="10.8416999999999994"/>
               </Node>
               <Node id="153" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="10.8416999999999994"/>
               </Node>
              </Node>
              <Node id="127">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="154" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="17"/>
               </Node>
               <Node id="155" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="17"/>
               </Node>
              </Node>
             </Node>
             <Node id="97">
              <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
              <Node id="128" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="129">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="156" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="25"/>
               </Node>
               <Node id="157" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="25"/>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="71" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
           <Node id="49">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="72" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="9.09999999999999964"/>
            </Node>
            <Node id="73" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="9.09999999999999964"/>
            </Node>
           </Node>
          </Node>
          <Node id="29">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="50">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="74">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="18.9291499999999999"/>
             <Node id="98">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="130" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="131" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
              </Node>
             </Node>
             <Node id="99" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="75" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="18.9291499999999999"/>
            </Node>
           </Node>
           <Node id="51">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="76" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="77">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             <Node id="100" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="101">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              <Node id="132" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="16"/>
              </Node>
              <Node id="133" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="16"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
          <Node id="30">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="15.25"/>
           <Node id="52">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="78">
             <SimplePredicate field="Age" operator="lessOrEqual" value="47.5"/>
             <Node id="102">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="134" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="11.2937499999999993"/>
              </Node>
              <Node id="135" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="11.2937499999999993"/>
              </Node>
             </Node>
             <Node id="103" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="79">
             <SimplePredicate field="Age" operator="greaterThan" value="47.5"/>
             <Node id="104" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="53.5"/>
             </Node>
             <Node id="105" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="53.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="53" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="31" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="15.25"/>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Fare" operator="greaterThan" value="25.0749999999999993"/>
         <Node id="18">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
          <Node id="32" score="1">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
          </Node>
          <Node id="33" score="1">
           <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
          </Node>
         </Node>
         <Node id="19" score="0">
          <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
        <Node id="10">
         <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
         <Node id="20">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="36.6875"/>
          <Node id="34" score="0">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
          </Node>
          <Node id="35" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Fare" operator="greaterThan" value="36.6875"/>
          <Node id="36" score="1">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="5.5"/>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="5.5"/>
          </Node>
         </Node>
        </Node>
        <Node id="11" score="1">
         <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;male&quot;</Array>
       </SimpleSetPredicate>
       <Node id="6">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="1" type="string">&quot;C&quot;</Array>
        </SimpleSetPredicate>
        <Node id="12" score="1">
         <SimplePredicate field="Age" operator="lessOrEqual" value="8"/>
        </Node>
        <Node id="13">
         <SimplePredicate field="Age" operator="greaterThan" value="8"/>
         <Node id="22">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
          <Node id="38">
           <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
           <Node id="54">
            <SimplePredicate field="Age" operator="lessOrEqual" value="22.5"/>
            <Node id="80">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="123.258299999999991"/>
             <Node id="106" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
             </Node>
             <Node id="107" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
             </Node>
            </Node>
            <Node id="81" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="123.258299999999991"/>
            </Node>
           </Node>
           <Node id="55" score="1">
            <SimplePredicate field="Age" operator="greaterThan" value="22.5"/>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
           <Node id="56" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="27.1354000000000006"/>
           </Node>
           <Node id="57">
            <SimplePredicate field="Fare" operator="greaterThan" value="27.1354000000000006"/>
            <Node id="82">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="108" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="28.7250000000000014"/>
             </Node>
             <Node id="109" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="28.7250000000000014"/>
             </Node>
            </Node>
            <Node id="83">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="110" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="111" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          <Node id="40">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           <Node id="58" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="59" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="41" score="0">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
        </SimpleSetPredicate>
        <Node id="14">
         <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
         <Node id="24">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="26.2687500000000007"/>
          <Node id="42">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="60">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
            <Node id="84">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
             <Node id="112">
              <SimplePredicate field="Age" operator="lessOrEqual" value="59.5"/>
              <Node id="136" score="0">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;Q&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="137">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="158" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="12.7624999999999993"/>
               </Node>
               <Node id="159">
                <SimplePredicate field="Fare" operator="greaterThan" value="12.7624999999999993"/>
                <Node id="166" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="29"/>
                </Node>
                <Node id="167" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="29"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="113">
              <SimplePredicate field="Age" operator="greaterThan" value="59.5"/>
              <Node id="138" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="64"/>
              </Node>
              <Node id="139" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="64"/>
              </Node>
             </Node>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
            </Node>
           </Node>
           <Node id="61">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            <Node id="86" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="16.5"/>
            </Node>
            <Node id="87">
             <SimplePredicate field="Age" operator="greaterThan" value="16.5"/>
             <Node id="114" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="28.75"/>
             </Node>
             <Node id="115">
              <SimplePredicate field="Age" operator="greaterThan" value="28.75"/>
              <Node id="140" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="9.25"/>
              </Node>
              <Node id="141" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="9.25"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="43">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="62">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="15.6999999999999993"/>
            <Node id="88">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="116" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.51250000000000018"/>
             </Node>
             <Node id="117">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.51250000000000018"/>
              <Node id="142" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="20.5"/>
              </Node>
              <Node id="143">
               <SimplePredicate field="Age" operator="greaterThan" value="20.5"/>
               <Node id="160" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
               </Node>
               <Node id="161" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="89" score="0">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="63">
            <SimplePredicate field="Fare" operator="greaterThan" value="15.6999999999999993"/>
            <Node id="90">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="118" score="0">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
             </Node>
             <Node id="119">
              <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
              <Node id="144" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="21.6083500000000015"/>
              </Node>
              <Node id="145" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="21.6083500000000015"/>
              </Node>
             </Node>
            </Node>
            <Node id="91">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="120" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="19.4812499999999993"/>
             </Node>
             <Node id="121">
              <SimplePredicate field="Fare" operator="greaterThan" value="19.4812499999999993"/>
              <Node id="146" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="10.5"/>
              </Node>
              <Node id="147" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="10.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="25">
          <SimplePredicate field="Fare" operator="greaterThan" value="26.2687500000000007"/>
          <Node id="44">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="64" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
           </Node>
           <Node id="65">
            <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
            <Node id="92" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="26.46875"/>
            </Node>
            <Node id="93">
             <SimplePredicate field="Fare" operator="greaterThan" value="26.46875"/>
             <Node id="122">
              <SimplePredicate field="Age" operator="lessOrEqual" value="72.5"/>
              <Node id="148">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="162">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="2"/>
                <Node id="168">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
                 <Node id="172">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="37"/>
                  <Node id="174" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="35.25"/>
                  </Node>
                  <Node id="175" score="1">
                   <SimplePredicate field="Fare" operator="greaterThan" value="35.25"/>
                  </Node>
                 </Node>
                 <Node id="173" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="37"/>
                 </Node>
                </Node>
                <Node id="169" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="53"/>
                </Node>
               </Node>
               <Node id="163">
                <SimplePredicate field="Pclass" operator="greaterThan" value="2"/>
                <Node id="170" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
                </Node>
                <Node id="171" score="1">
                 <SimplePredicate field="Age" operator="greaterThan" value="30"/>
                </Node>
               </Node>
              </Node>
              <Node id="149" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="123" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="72.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="45">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="66">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
            <Node id="94">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="115.200000000000003"/>
             <Node id="124" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="12"/>
             </Node>
             <Node id="125">
              <SimplePredicate field="Age" operator="greaterThan" value="12"/>
              <Node id="150">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="164" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="33.5"/>
               </Node>
               <Node id="165" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="33.5"/>
               </Node>
              </Node>
              <Node id="151" score="0">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
            <Node id="95" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="115.200000000000003"/>
            </Node>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15">
         <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
         <Node id="26">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="35.5375000000000014"/>
          <Node id="46">
           <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
           <Node id="68" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
           </Node>
           <Node id="69" score="1">
            <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
           </Node>
          </Node>
          <Node id="47" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
         <Node id="27" score="0">
          <SimplePredicate field="Fare" operator="greaterThan" value="35.5375000000000014"/>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="9">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
       <Node id="4">
        <SimplePredicate field="Age" operator="lessOrEqual" value="16.5"/>
        <Node id="8">
         <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
         <Node id="16" score="1">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
          </SimpleSetPredicate>
         </Node>
         <Node id="17">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="28" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="56.25"/>
          </Node>
          <Node id="29" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="56.25"/>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
         <Node id="18">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="25.1999999999999993"/>
          <Node id="30" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="9.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Age" operator="greaterThan" value="9.5"/>
           <Node id="46" score="1">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="47" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="19" score="1">
          <SimplePredicate field="Fare" operator="greaterThan" value="25.1999999999999993"/>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Age" operator="greaterThan" value="16.5"/>
        <Node id="10">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;&quot;   &quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="32">
           <SimplePredicate field="Age" operator="lessOrEqual" value="20"/>
           <Node id="48">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="68" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="13.2291500000000006"/>
            </Node>
            <Node id="69" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="13.2291500000000006"/>
            </Node>
           </Node>
           <Node id="49" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="33">
           <SimplePredicate field="Age" operator="greaterThan" value="20"/>
           <Node id="50">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="70" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="47"/>
            </Node>
            <Node id="71">
             <SimplePredicate field="Age" operator="greaterThan" value="47"/>
             <Node id="98">
              <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
              <Node id="130" score="0">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="131" score="1">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="99" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="53"/>
             </Node>
            </Node>
           </Node>
           <Node id="51">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="72" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2"/>
            </Node>
            <Node id="73">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2"/>
             <Node id="100" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
             </Node>
             <Node id="101" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="34">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
           <Node id="52">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="74">
             <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
             <Node id="102" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="9.83960000000000079"/>
             </Node>
             <Node id="103">
              <SimplePredicate field="Fare" operator="greaterThan" value="9.83960000000000079"/>
              <Node id="132" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;Q&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="133" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
              </Node>
             </Node>
            </Node>
            <Node id="75" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="39"/>
            </Node>
           </Node>
           <Node id="53">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            <Node id="76" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="77">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             <Node id="104" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="34"/>
             </Node>
             <Node id="105" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="34"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="35" score="0">
           <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22">
          <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
          <Node id="36">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           <Node id="54">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
            <Node id="78" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
            </Node>
            <Node id="79">
             <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
             <Node id="106" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="28.7250000000000014"/>
             </Node>
             <Node id="107">
              <SimplePredicate field="Fare" operator="greaterThan" value="28.7250000000000014"/>
              <Node id="134">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="38.9500000000000028"/>
               <Node id="150">
                <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                <Node id="158" score="0">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;C&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
                <Node id="159" score="1">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
               </Node>
               <Node id="151" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
               </Node>
              </Node>
              <Node id="135" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="38.9500000000000028"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
          <Node id="38">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="26.1437500000000007"/>
           <Node id="56" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="57">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            <Node id="80">
             <SimplePredicate field="Age" operator="lessOrEqual" value="41.5"/>
             <Node id="108" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="12.9375"/>
             </Node>
             <Node id="109">
              <SimplePredicate field="Fare" operator="greaterThan" value="12.9375"/>
              <Node id="136" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="13.75"/>
              </Node>
              <Node id="137" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="13.75"/>
              </Node>
             </Node>
            </Node>
            <Node id="81">
             <SimplePredicate field="Age" operator="greaterThan" value="41.5"/>
             <Node id="110" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;Q&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="111">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
              <Node id="138" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.84999999999999964"/>
              </Node>
              <Node id="139">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.84999999999999964"/>
               <Node id="152">
                <SimplePredicate field="Age" operator="lessOrEqual" value="45.5"/>
                <Node id="160" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="7.98750000000000071"/>
                </Node>
                <Node id="161">
                 <SimplePredicate field="Fare" operator="greaterThan" value="7.98750000000000071"/>
                 <Node id="168" score="0">
                  <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
                 </Node>
                 <Node id="169">
                  <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                  <Node id="174" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="8.35624999999999929"/>
                  </Node>
                  <Node id="175" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="8.35624999999999929"/>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="153">
                <SimplePredicate field="Age" operator="greaterThan" value="45.5"/>
                <Node id="162" score="0">
                 <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
                </Node>
                <Node id="163" score="0">
                 <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Fare" operator="greaterThan" value="26.1437500000000007"/>
           <Node id="58">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="82">
             <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
             <Node id="112" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="35.3125"/>
             </Node>
             <Node id="113" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="35.3125"/>
             </Node>
            </Node>
            <Node id="83">
             <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
             <Node id="114" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="57"/>
             </Node>
             <Node id="115" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="57"/>
             </Node>
            </Node>
           </Node>
           <Node id="59">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="84" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="53"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
       <Node id="6">
        <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
        <Node id="12">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="24">
          <SimplePredicate field="Age" operator="lessOrEqual" value="48.5"/>
          <Node id="40">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           <Node id="60">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="86" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="22"/>
            </Node>
            <Node id="87" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="22"/>
            </Node>
           </Node>
           <Node id="61">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="88" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="89">
             <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
             <Node id="116" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="135.775000000000006"/>
             </Node>
             <Node id="117">
              <SimplePredicate field="Fare" operator="greaterThan" value="135.775000000000006"/>
              <Node id="140" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="207.275000000000006"/>
              </Node>
              <Node id="141" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="207.275000000000006"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="41">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           <Node id="62">
            <SimplePredicate field="Age" operator="lessOrEqual" value="39"/>
            <Node id="90">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
             <Node id="118">
              <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
              <Node id="142" score="1">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;Q&quot;</Array>
               </SimpleSetPredicate>
              </Node>
              <Node id="143">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="154">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                <Node id="164">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
                 <Node id="170" score="0">
                  <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                 </Node>
                 <Node id="171">
                  <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                  <Node id="176" score="1">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="6.5"/>
                  </Node>
                  <Node id="177" score="0">
                   <SimplePredicate field="Age" operator="greaterThan" value="6.5"/>
                  </Node>
                 </Node>
                </Node>
                <Node id="165" score="1">
                 <SimplePredicate field="Age" operator="greaterThan" value="27"/>
                </Node>
               </Node>
               <Node id="155" score="0">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="119">
              <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
              <Node id="144">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
               <Node id="156" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;Q&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="157">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="166" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
                </Node>
                <Node id="167">
                 <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
                 <Node id="172" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="33"/>
                 </Node>
                 <Node id="173" score="1">
                  <SimplePredicate field="Age" operator="greaterThan" value="33"/>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="145" score="1">
               <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
              </Node>
             </Node>
            </Node>
            <Node id="91">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
             <Node id="120" score="1">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="121" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="63" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="39"/>
           </Node>
          </Node>
         </Node>
         <Node id="25" score="1">
          <SimplePredicate field="Age" operator="greaterThan" value="48.5"/>
         </Node>
        </Node>
        <Node id="13">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="26">
          <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
          <Node id="42" score="1">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
          </Node>
          <Node id="43" score="1">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
         <Node id="27">
          <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
          <Node id="44">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="52.2771000000000043"/>
           <Node id="64">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="92">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
             <Node id="122" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="123" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
            <Node id="93" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
           <Node id="65">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="94">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
             <Node id="124" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="12"/>
             </Node>
             <Node id="125" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="12"/>
             </Node>
            </Node>
            <Node id="95" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="45">
           <SimplePredicate field="Fare" operator="greaterThan" value="52.2771000000000043"/>
           <Node id="66" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="23"/>
           </Node>
           <Node id="67">
            <SimplePredicate field="Age" operator="greaterThan" value="23"/>
            <Node id="96">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;C&quot;</Array>
             </SimpleSetPredicate>
             <Node id="126">
              <SimplePredicate field="Age" operator="lessOrEqual" value="49.5"/>
              <Node id="146" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="93.8062500000000057"/>
              </Node>
              <Node id="147" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="93.8062500000000057"/>
              </Node>
             </Node>
             <Node id="127" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="49.5"/>
             </Node>
            </Node>
            <Node id="97">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="128" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="28"/>
             </Node>
             <Node id="129">
              <SimplePredicate field="Age" operator="greaterThan" value="28"/>
              <Node id="148" score="0">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="149" score="0">
               <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
        <Node id="14" score="0">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
        </Node>
        <Node id="15" score="0">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="10">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Fare" operator="lessOrEqual" value="48.2999999999999972"/>
       <Node id="4">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;female&quot;</Array>
        </SimpleSetPredicate>
        <Node id="8">
         <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
         <Node id="14">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="25.9646000000000008"/>
          <Node id="24">
           <SimplePredicate field="Age" operator="lessOrEqual" value="55.5"/>
           <Node id="42" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
           </Node>
           <Node id="43" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
           </Node>
          </Node>
          <Node id="25" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="55.5"/>
          </Node>
         </Node>
         <Node id="15">
          <SimplePredicate field="Fare" operator="greaterThan" value="25.9646000000000008"/>
          <Node id="26">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;C&quot;</Array>
           </SimpleSetPredicate>
           <Node id="44">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            <Node id="70" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="47"/>
            </Node>
            <Node id="71" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="47"/>
            </Node>
           </Node>
           <Node id="45" score="1">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
          <Node id="27">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="46">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="72">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="27.5"/>
             <Node id="96">
              <SimplePredicate field="Age" operator="lessOrEqual" value="43"/>
              <Node id="116">
               <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
               <Node id="132" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="25"/>
               </Node>
               <Node id="133" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="25"/>
               </Node>
              </Node>
              <Node id="117" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
              </Node>
             </Node>
             <Node id="97">
              <SimplePredicate field="Age" operator="greaterThan" value="43"/>
              <Node id="118" score="0">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="119" score="1">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
            </Node>
            <Node id="73" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="27.5"/>
            </Node>
           </Node>
           <Node id="47" score="1">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
         <Node id="16">
          <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
          <Node id="28">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
           <Node id="48">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="74" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="6.98749999999999982"/>
            </Node>
            <Node id="75">
             <SimplePredicate field="Fare" operator="greaterThan" value="6.98749999999999982"/>
             <Node id="98" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.74375000000000036"/>
             </Node>
             <Node id="99">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.74375000000000036"/>
              <Node id="120">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
               <Node id="134" score="0">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="135">
                <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                <Node id="146">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                 <Node id="152" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="11"/>
                 </Node>
                 <Node id="153" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="11"/>
                 </Node>
                </Node>
                <Node id="147">
                 <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                 <Node id="154">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                  <Node id="160" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="15.4937499999999986"/>
                  </Node>
                  <Node id="161" score="1">
                   <SimplePredicate field="Fare" operator="greaterThan" value="15.4937499999999986"/>
                  </Node>
                 </Node>
                 <Node id="155" score="1">
                  <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="121">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="136" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.78960000000000008"/>
               </Node>
               <Node id="137" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.78960000000000008"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="49" score="0">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
           </Node>
          </Node>
          <Node id="29" score="0">
           <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
          <Node id="30">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
           </SimpleSetPredicate>
           <Node id="50">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="15.3728999999999996"/>
            <Node id="76" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="14.75"/>
            </Node>
            <Node id="77" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="14.75"/>
            </Node>
           </Node>
           <Node id="51" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="15.3728999999999996"/>
           </Node>
          </Node>
          <Node id="31" score="0">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;male&quot;</Array>
        </SimpleSetPredicate>
        <Node id="10">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="1" type="string">&quot;C&quot;</Array>
         </SimpleSetPredicate>
         <Node id="18">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
          <Node id="32">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="28.7250000000000014"/>
           <Node id="52" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="27.1354000000000006"/>
           </Node>
           <Node id="53" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="27.1354000000000006"/>
           </Node>
          </Node>
          <Node id="33" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="28.7250000000000014"/>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          <Node id="34">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="54">
            <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
            <Node id="78">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="100" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="15.5"/>
             </Node>
             <Node id="101" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="15.5"/>
             </Node>
            </Node>
            <Node id="79" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="55">
            <SimplePredicate field="Age" operator="greaterThan" value="27"/>
            <Node id="80" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="13.3687500000000004"/>
            </Node>
            <Node id="81" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="13.3687500000000004"/>
            </Node>
           </Node>
          </Node>
          <Node id="35">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="56" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="57" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="20" score="0">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
         </Node>
         <Node id="21">
          <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
          <Node id="36">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           <Node id="58">
            <SimplePredicate field="Age" operator="lessOrEqual" value="6.5"/>
            <Node id="82" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="83" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
           <Node id="59">
            <SimplePredicate field="Age" operator="greaterThan" value="6.5"/>
            <Node id="84">
             <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
             <Node id="102" score="0">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
             </Node>
             <Node id="103" score="0">
              <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             </Node>
            </Node>
            <Node id="85">
             <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
             <Node id="104">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;Q&quot;</Array>
              </SimpleSetPredicate>
              <Node id="122">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="19.375"/>
               <Node id="138" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.78960000000000008"/>
               </Node>
               <Node id="139" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.78960000000000008"/>
               </Node>
              </Node>
              <Node id="123" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="19.375"/>
              </Node>
             </Node>
             <Node id="105" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="37">
           <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           <Node id="60">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
            <Node id="86" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="20.5499999999999972"/>
            </Node>
            <Node id="87">
             <SimplePredicate field="Fare" operator="greaterThan" value="20.5499999999999972"/>
             <Node id="106">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="124" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="27.5"/>
              </Node>
              <Node id="125" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="27.5"/>
              </Node>
             </Node>
             <Node id="107">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
              <Node id="126">
               <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
               <Node id="140" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="2"/>
               </Node>
               <Node id="141" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="2"/>
               </Node>
              </Node>
              <Node id="127" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Fare" operator="greaterThan" value="48.2999999999999972"/>
       <Node id="6">
        <SimplePredicate field="Age" operator="lessOrEqual" value="62"/>
        <Node id="12">
         <SimplePredicate field="SibSp" operator="lessOrEqual" value="5.5"/>
         <Node id="22">
          <SimplePredicate field="Age" operator="lessOrEqual" value="21.5"/>
          <Node id="38">
           <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
           <Node id="62" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="63" score="1">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
           <Node id="64">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            <Node id="88" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="77.6228999999999871"/>
            </Node>
            <Node id="89">
             <SimplePredicate field="Fare" operator="greaterThan" value="77.6228999999999871"/>
             <Node id="108" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="94.2750000000000057"/>
             </Node>
             <Node id="109">
              <SimplePredicate field="Fare" operator="greaterThan" value="94.2750000000000057"/>
              <Node id="128" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="168.212500000000006"/>
              </Node>
              <Node id="129" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="168.212500000000006"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="65" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Age" operator="greaterThan" value="21.5"/>
          <Node id="40">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="66" score="1">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="67">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="90" score="1">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;C&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="91" score="1">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
          </Node>
          <Node id="41">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
           <Node id="68">
            <SimplePredicate field="Age" operator="lessOrEqual" value="49.5"/>
            <Node id="92">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
             <Node id="110">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
              <Node id="130">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="77.7896000000000072"/>
               <Node id="142" score="1">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="143">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="148" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                </Node>
                <Node id="149">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 <Node id="156" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="29"/>
                 </Node>
                 <Node id="157">
                  <SimplePredicate field="Age" operator="greaterThan" value="29"/>
                  <Node id="162" score="0">
                   <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                  </Node>
                  <Node id="163" score="1">
                   <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                  </Node>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="131">
               <SimplePredicate field="Fare" operator="greaterThan" value="77.7896000000000072"/>
               <Node id="144" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="85.6375000000000028"/>
               </Node>
               <Node id="145">
                <SimplePredicate field="Fare" operator="greaterThan" value="85.6375000000000028"/>
                <Node id="150">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 <Node id="158" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="369.927099999999996"/>
                 </Node>
                 <Node id="159" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="369.927099999999996"/>
                 </Node>
                </Node>
                <Node id="151" score="1">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="111" score="0">
              <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
            <Node id="93" score="1">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
           <Node id="69">
            <SimplePredicate field="Age" operator="greaterThan" value="49.5"/>
            <Node id="94">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;C&quot;</Array>
             </SimpleSetPredicate>
             <Node id="112" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="55.5"/>
             </Node>
             <Node id="113" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="55.5"/>
             </Node>
            </Node>
            <Node id="95">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="114" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="106.650000000000006"/>
             </Node>
             <Node id="115" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="106.650000000000006"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="13" score="0">
         <SimplePredicate field="SibSp" operator="greaterThan" value="5.5"/>
        </Node>
       </Node>
       <Node id="7" score="0">
        <SimplePredicate field="Age" operator="greaterThan" value="62"/>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="11">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Fare" operator="lessOrEqual" value="10.4812499999999993"/>
       <Node id="4">
        <SimplePredicate field="Age" operator="lessOrEqual" value="32.5"/>
        <Node id="8">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="1" type="string">&quot;Q&quot;</Array>
         </SimpleSetPredicate>
         <Node id="16">
          <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
          <Node id="30" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.74375000000000036"/>
          </Node>
          <Node id="31" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.74375000000000036"/>
          </Node>
         </Node>
         <Node id="17" score="0">
          <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
         </Node>
        </Node>
        <Node id="9">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="18">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="32">
           <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
           <Node id="48">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="7.875"/>
            <Node id="70" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
            </Node>
            <Node id="71" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
            </Node>
           </Node>
           <Node id="49">
            <SimplePredicate field="Fare" operator="greaterThan" value="7.875"/>
            <Node id="72">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
             <Node id="92" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="93" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="73" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
            </Node>
           </Node>
          </Node>
          <Node id="33">
           <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
           <Node id="50" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="51" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="19">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
          <Node id="34">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="9.49164999999999992"/>
           <Node id="52">
            <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
            <Node id="74">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="94">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.79790000000000028"/>
              <Node id="116" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.7854000000000001"/>
              </Node>
              <Node id="117" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.7854000000000001"/>
              </Node>
             </Node>
             <Node id="95" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.79790000000000028"/>
             </Node>
            </Node>
            <Node id="75">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="96" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="97">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
              <Node id="118" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
              </Node>
              <Node id="119" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="53" score="1">
            <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
           </Node>
          </Node>
          <Node id="35">
           <SimplePredicate field="Fare" operator="greaterThan" value="9.49164999999999992"/>
           <Node id="54" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
           </Node>
           <Node id="55" score="1">
            <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimplePredicate field="Age" operator="greaterThan" value="32.5"/>
        <Node id="10" score="0">
         <SimplePredicate field="Pclass" operator="lessOrEqual" value="2"/>
        </Node>
        <Node id="11">
         <SimplePredicate field="Pclass" operator="greaterThan" value="2"/>
         <Node id="20" score="0">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
          </SimpleSetPredicate>
         </Node>
         <Node id="21">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="36" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.97290000000000099"/>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.97290000000000099"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Fare" operator="greaterThan" value="10.4812499999999993"/>
       <Node id="6">
        <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
        <Node id="12">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;&quot;   &quot;C&quot;</Array>
         </SimpleSetPredicate>
         <Node id="22" score="1">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;&quot;</Array>
          </SimpleSetPredicate>
         </Node>
         <Node id="23">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;C&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="38" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="29.8500000000000014"/>
          </Node>
          <Node id="39">
           <SimplePredicate field="Fare" operator="greaterThan" value="29.8500000000000014"/>
           <Node id="56">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="62.6687500000000028"/>
            <Node id="76">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="98">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              <Node id="120" score="1">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              </Node>
              <Node id="121" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="99" score="1">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
            <Node id="77" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="57">
            <SimplePredicate field="Fare" operator="greaterThan" value="62.6687500000000028"/>
            <Node id="78" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="79">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="100">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              <Node id="122">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="98.7520999999999987"/>
               <Node id="142" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="35.5"/>
               </Node>
               <Node id="143">
                <SimplePredicate field="Age" operator="greaterThan" value="35.5"/>
                <Node id="164" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                </Node>
                <Node id="165" score="1">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="123">
               <SimplePredicate field="Fare" operator="greaterThan" value="98.7520999999999987"/>
               <Node id="144">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                <Node id="166" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="379.925000000000011"/>
                </Node>
                <Node id="167" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="379.925000000000011"/>
                </Node>
               </Node>
               <Node id="145" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="101" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
          <Array n="2" type="string">&quot;Q&quot;   &quot;S&quot;</Array>
         </SimpleSetPredicate>
         <Node id="24">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
          <Node id="40" score="0">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;Q&quot;</Array>
           </SimpleSetPredicate>
          </Node>
          <Node id="41">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="58">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="79.25"/>
            <Node id="80" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="25.9271000000000029"/>
            </Node>
            <Node id="81">
             <SimplePredicate field="Fare" operator="greaterThan" value="25.9271000000000029"/>
             <Node id="102" score="1">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="103">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
              <Node id="124">
               <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               <Node id="146" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="31.4103999999999992"/>
               </Node>
               <Node id="147" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="31.4103999999999992"/>
               </Node>
              </Node>
              <Node id="125">
               <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               <Node id="148" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="52.8271000000000015"/>
               </Node>
               <Node id="149">
                <SimplePredicate field="Fare" operator="greaterThan" value="52.8271000000000015"/>
                <Node id="168" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="67.9249999999999972"/>
                </Node>
                <Node id="169" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="67.9249999999999972"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="59">
            <SimplePredicate field="Fare" operator="greaterThan" value="79.25"/>
            <Node id="82" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="83">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="104">
              <SimplePredicate field="Age" operator="lessOrEqual" value="42"/>
              <Node id="126">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;female&quot;</Array>
               </SimpleSetPredicate>
               <Node id="150" score="1">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
               </Node>
               <Node id="151">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                <Node id="170" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
                </Node>
                <Node id="171" score="1">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
                </Node>
               </Node>
              </Node>
              <Node id="127">
               <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                <Array n="1" type="string">&quot;male&quot;</Array>
               </SimpleSetPredicate>
               <Node id="152" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="15"/>
               </Node>
               <Node id="153">
                <SimplePredicate field="Age" operator="greaterThan" value="15"/>
                <Node id="172" score="1">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="1"/>
                </Node>
                <Node id="173" score="0">
                 <SimplePredicate field="Parch" operator="greaterThan" value="1"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="105">
              <SimplePredicate field="Age" operator="greaterThan" value="42"/>
              <Node id="128" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="88.4874999999999972"/>
              </Node>
              <Node id="129" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="88.4874999999999972"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="25" score="0">
          <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
        <Node id="14">
         <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
         <Node id="26" score="1">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="20.8249999999999993"/>
         </Node>
         <Node id="27">
          <SimplePredicate field="Fare" operator="greaterThan" value="20.8249999999999993"/>
          <Node id="42">
           <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
           <Node id="60" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
          <Node id="43">
           <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
           <Node id="62" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="63" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15">
         <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
         <Node id="28">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
          <Node id="44">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           <Node id="64" score="1">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="65">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="84">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
             <Node id="106">
              <SimplePredicate field="Age" operator="lessOrEqual" value="55.5"/>
              <Node id="130" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="26.125"/>
              </Node>
              <Node id="131" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="26.125"/>
              </Node>
             </Node>
             <Node id="107" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="55.5"/>
             </Node>
            </Node>
            <Node id="85" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="45">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           <Node id="66" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="12"/>
           </Node>
           <Node id="67">
            <SimplePredicate field="Age" operator="greaterThan" value="12"/>
            <Node id="86">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="24.8083499999999972"/>
             <Node id="108">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
              <Node id="132" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="13.4375"/>
              </Node>
              <Node id="133">
               <SimplePredicate field="Fare" operator="greaterThan" value="13.4375"/>
               <Node id="154">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                <Node id="174">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                 <Node id="178" score="0">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="21"/>
                 </Node>
                 <Node id="179">
                  <SimplePredicate field="Age" operator="greaterThan" value="21"/>
                  <Node id="182" score="1">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;Q&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="183" score="0">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                 </Node>
                </Node>
                <Node id="175">
                 <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                 <Node id="180" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="17.3249999999999993"/>
                 </Node>
                 <Node id="181" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="17.3249999999999993"/>
                 </Node>
                </Node>
               </Node>
               <Node id="155">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                <Node id="176" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="22.9041500000000013"/>
                </Node>
                <Node id="177" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="22.9041500000000013"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="109" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
             </Node>
            </Node>
            <Node id="87">
             <SimplePredicate field="Fare" operator="greaterThan" value="24.8083499999999972"/>
             <Node id="110">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="32.8812500000000014"/>
              <Node id="134">
               <SimplePredicate field="Age" operator="lessOrEqual" value="38.5"/>
               <Node id="156" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="28.4270999999999994"/>
               </Node>
               <Node id="157" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="28.4270999999999994"/>
               </Node>
              </Node>
              <Node id="135" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="38.5"/>
              </Node>
             </Node>
             <Node id="111" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="32.8812500000000014"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="29">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
          <Node id="46">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
           <Node id="68">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="88">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="112">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="136" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
              </Node>
              <Node id="137">
               <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
               <Node id="158" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="13.3687500000000004"/>
               </Node>
               <Node id="159" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="13.3687500000000004"/>
               </Node>
              </Node>
             </Node>
             <Node id="113" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="89" score="1">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="69">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="90" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="12.5"/>
            </Node>
            <Node id="91">
             <SimplePredicate field="Age" operator="greaterThan" value="12.5"/>
             <Node id="114">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="47.7479000000000013"/>
              <Node id="138">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="160" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="11"/>
               </Node>
               <Node id="161" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="11"/>
               </Node>
              </Node>
              <Node id="139" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="115">
              <SimplePredicate field="Fare" operator="greaterThan" value="47.7479000000000013"/>
              <Node id="140" score="0">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
              </Node>
              <Node id="141">
               <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
               <Node id="162" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
               </Node>
               <Node id="163" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="30"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="47" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="12">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Fare" operator="lessOrEqual" value="9.48750000000000071"/>
       <Node id="4" score="1">
        <SimplePredicate field="Age" operator="lessOrEqual" value="14"/>
       </Node>
       <Node id="5">
        <SimplePredicate field="Age" operator="greaterThan" value="14"/>
        <Node id="8">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="7.13335000000000008"/>
         <Node id="14" score="0">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="3.24790000000000001"/>
         </Node>
         <Node id="15" score="0">
          <SimplePredicate field="Fare" operator="greaterThan" value="3.24790000000000001"/>
         </Node>
        </Node>
        <Node id="9">
         <SimplePredicate field="Fare" operator="greaterThan" value="7.13335000000000008"/>
         <Node id="16">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;Q&quot;</Array>
          </SimpleSetPredicate>
          <Node id="26">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.76874999999999982"/>
           <Node id="44" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="7.73124999999999929"/>
           </Node>
           <Node id="45">
            <SimplePredicate field="Fare" operator="greaterThan" value="7.73124999999999929"/>
            <Node id="72" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="73" score="0">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
          </Node>
          <Node id="27">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.76874999999999982"/>
           <Node id="46" score="1">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="47" score="0">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
           </Node>
          </Node>
         </Node>
         <Node id="17">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="28" score="1">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.1833499999999999"/>
          </Node>
          <Node id="29">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.1833499999999999"/>
           <Node id="48">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="74">
             <SimplePredicate field="Age" operator="lessOrEqual" value="22.75"/>
             <Node id="104">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="134" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.22710000000000008"/>
              </Node>
              <Node id="135" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.22710000000000008"/>
              </Node>
             </Node>
             <Node id="105" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="75">
             <SimplePredicate field="Age" operator="greaterThan" value="22.75"/>
             <Node id="106">
              <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
              <Node id="136">
               <SimplePredicate field="Age" operator="lessOrEqual" value="28.75"/>
               <Node id="160" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.22710000000000008"/>
               </Node>
               <Node id="161" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.22710000000000008"/>
               </Node>
              </Node>
              <Node id="137" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="28.75"/>
              </Node>
             </Node>
             <Node id="107" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="49">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="76">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="108">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;female&quot;</Array>
              </SimpleSetPredicate>
              <Node id="138">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.98750000000000071"/>
               <Node id="162" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.52289999999999992"/>
               </Node>
               <Node id="163">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.52289999999999992"/>
                <Node id="186">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="7.81460000000000043"/>
                 <Node id="198">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
                  <Node id="210" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="19.5"/>
                  </Node>
                  <Node id="211" score="1">
                   <SimplePredicate field="Age" operator="greaterThan" value="19.5"/>
                  </Node>
                 </Node>
                 <Node id="199" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
                 </Node>
                </Node>
                <Node id="187" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="7.81460000000000043"/>
                </Node>
               </Node>
              </Node>
              <Node id="139">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.98750000000000071"/>
               <Node id="164">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                <Node id="188" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="8.6728999999999985"/>
                </Node>
                <Node id="189" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="8.6728999999999985"/>
                </Node>
               </Node>
               <Node id="165" score="0">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="109">
              <SimpleSetPredicate field="Sex" booleanOperator="isIn">
               <Array n="1" type="string">&quot;male&quot;</Array>
              </SimpleSetPredicate>
              <Node id="140">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.9104000000000001"/>
               <Node id="166" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="7.82499999999999929"/>
               </Node>
               <Node id="167" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="7.82499999999999929"/>
               </Node>
              </Node>
              <Node id="141">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.9104000000000001"/>
               <Node id="168" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="17.5"/>
               </Node>
               <Node id="169">
                <SimplePredicate field="Age" operator="greaterThan" value="17.5"/>
                <Node id="190">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
                 <Node id="200" score="0">
                  <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 </Node>
                 <Node id="201" score="1">
                  <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                 </Node>
                </Node>
                <Node id="191" score="0">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="77">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="110" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="8.87709999999999866"/>
             </Node>
             <Node id="111" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="8.87709999999999866"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Fare" operator="greaterThan" value="9.48750000000000071"/>
       <Node id="6">
        <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
        <Node id="10">
         <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
         <Node id="18">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="77.6228999999999871"/>
          <Node id="30">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           <Node id="50">
            <SimplePredicate field="Age" operator="lessOrEqual" value="53"/>
            <Node id="78">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="112">
              <SimplePredicate field="Age" operator="lessOrEqual" value="47.5"/>
              <Node id="142" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
              </Node>
              <Node id="143">
               <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
               <Node id="170">
                <SimplePredicate field="Age" operator="lessOrEqual" value="45.25"/>
                <Node id="192">
                 <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;female&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="202">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
                  <Node id="212">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
                   <Node id="220" score="1">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8249999999999993"/>
                   </Node>
                   <Node id="221">
                    <SimplePredicate field="Fare" operator="greaterThan" value="12.8249999999999993"/>
                    <Node id="226" score="1">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="26"/>
                    </Node>
                    <Node id="227" score="1">
                     <SimplePredicate field="Age" operator="greaterThan" value="26"/>
                    </Node>
                   </Node>
                  </Node>
                  <Node id="213" score="0">
                   <SimplePredicate field="Age" operator="greaterThan" value="37"/>
                  </Node>
                 </Node>
                 <Node id="203" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
                 </Node>
                </Node>
                <Node id="193">
                 <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;male&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="204">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="25.5"/>
                  <Node id="214" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="21"/>
                  </Node>
                  <Node id="215" score="0">
                   <SimplePredicate field="Age" operator="greaterThan" value="21"/>
                  </Node>
                 </Node>
                 <Node id="205">
                  <SimplePredicate field="Age" operator="greaterThan" value="25.5"/>
                  <Node id="216">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
                   <Node id="222" score="0">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="28"/>
                   </Node>
                   <Node id="223" score="1">
                    <SimplePredicate field="Fare" operator="greaterThan" value="28"/>
                   </Node>
                  </Node>
                  <Node id="217">
                   <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
                   <Node id="224">
                    <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                     <Array n="1" type="string">&quot;C&quot;</Array>
                    </SimpleSetPredicate>
                    <Node id="228" score="0">
                     <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                    </Node>
                    <Node id="229" score="0">
                     <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
                    </Node>
                   </Node>
                   <Node id="225">
                    <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                     <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                    </SimpleSetPredicate>
                    <Node id="230" score="0">
                     <SimplePredicate field="Fare" operator="lessOrEqual" value="26.1437500000000007"/>
                    </Node>
                    <Node id="231">
                     <SimplePredicate field="Fare" operator="greaterThan" value="26.1437500000000007"/>
                     <Node id="232" score="0">
                      <SimplePredicate field="Age" operator="lessOrEqual" value="33"/>
                     </Node>
                     <Node id="233" score="1">
                      <SimplePredicate field="Age" operator="greaterThan" value="33"/>
                     </Node>
                    </Node>
                   </Node>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="171" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="45.25"/>
               </Node>
              </Node>
             </Node>
             <Node id="113">
              <SimplePredicate field="Age" operator="greaterThan" value="47.5"/>
              <Node id="144" score="1">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="145">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
               <Node id="172" score="1">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;female&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="173" score="0">
                <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;male&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="79">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="114" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="59.6791999999999945"/>
             </Node>
             <Node id="115" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="59.6791999999999945"/>
             </Node>
            </Node>
           </Node>
           <Node id="51">
            <SimplePredicate field="Age" operator="greaterThan" value="53"/>
            <Node id="80" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="81">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="116">
              <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
              <Node id="146" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="35.0771000000000015"/>
              </Node>
              <Node id="147">
               <SimplePredicate field="Fare" operator="greaterThan" value="35.0771000000000015"/>
               <Node id="174" score="0">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="175" score="0">
                <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="117" score="0">
              <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="31">
           <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           <Node id="52" score="1">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="53">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="82" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="83" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Fare" operator="greaterThan" value="77.6228999999999871"/>
          <Node id="32" score="1">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
          </Node>
          <Node id="33">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
           <Node id="54" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="14"/>
           </Node>
           <Node id="55">
            <SimplePredicate field="Age" operator="greaterThan" value="14"/>
            <Node id="84">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             <Node id="118" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
             </Node>
             <Node id="119">
              <SimplePredicate field="Age" operator="greaterThan" value="32"/>
              <Node id="148" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="332.895849999999996"/>
              </Node>
              <Node id="149" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="332.895849999999996"/>
              </Node>
             </Node>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
         <Node id="20" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
         </Node>
         <Node id="21">
          <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
          <Node id="34">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;C&quot;</Array>
           </SimpleSetPredicate>
           <Node id="56">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="50.4895999999999958"/>
            <Node id="86" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="35.8250000000000028"/>
            </Node>
            <Node id="87" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="35.8250000000000028"/>
            </Node>
           </Node>
           <Node id="57">
            <SimplePredicate field="Fare" operator="greaterThan" value="50.4895999999999958"/>
            <Node id="88" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="89">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="120">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="98.7520999999999987"/>
              <Node id="150" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="38"/>
              </Node>
              <Node id="151" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="38"/>
              </Node>
             </Node>
             <Node id="121" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="98.7520999999999987"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="35">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="58">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="90" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="38.5"/>
            </Node>
            <Node id="91" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="38.5"/>
            </Node>
           </Node>
           <Node id="59">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="92" score="1">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;female&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="93">
             <SimpleSetPredicate field="Sex" booleanOperator="isIn">
              <Array n="1" type="string">&quot;male&quot;</Array>
             </SimpleSetPredicate>
             <Node id="122">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="45.5"/>
              <Node id="152">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="26.125"/>
               <Node id="176" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="33"/>
               </Node>
               <Node id="177" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="33"/>
               </Node>
              </Node>
              <Node id="153" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="26.125"/>
              </Node>
             </Node>
             <Node id="123">
              <SimplePredicate field="Fare" operator="greaterThan" value="45.5"/>
              <Node id="154">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
               <Node id="178">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="54.5"/>
                <Node id="194">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
                 <Node id="206" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
                 </Node>
                 <Node id="207" score="1">
                  <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
                 </Node>
                </Node>
                <Node id="195" score="1">
                 <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
               <Node id="179" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="54.5"/>
               </Node>
              </Node>
              <Node id="155" score="0">
               <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
        <Node id="12">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="13.9083500000000004"/>
         <Node id="22">
          <SimplePredicate field="Age" operator="lessOrEqual" value="18.5"/>
          <Node id="36">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="10.7979000000000003"/>
           <Node id="60" score="1">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="10.1521000000000008"/>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="10.1521000000000008"/>
           </Node>
          </Node>
          <Node id="37" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="10.7979000000000003"/>
          </Node>
         </Node>
         <Node id="23">
          <SimplePredicate field="Age" operator="greaterThan" value="18.5"/>
          <Node id="38">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
           <Node id="62" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="29"/>
           </Node>
           <Node id="63">
            <SimplePredicate field="Age" operator="greaterThan" value="29"/>
            <Node id="94" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="31.5"/>
            </Node>
            <Node id="95" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="31.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="39" score="1">
           <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimplePredicate field="Fare" operator="greaterThan" value="13.9083500000000004"/>
         <Node id="24">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="40">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="64">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="20.625"/>
            <Node id="96">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
             <Node id="124" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="125" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
            <Node id="97" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="27"/>
            </Node>
           </Node>
           <Node id="65" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="20.625"/>
           </Node>
          </Node>
          <Node id="41">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
           <Node id="66">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="98" score="0">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;Q&quot;</Array>
             </SimpleSetPredicate>
            </Node>
            <Node id="99">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
             <Node id="126" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="40.3229000000000042"/>
             </Node>
             <Node id="127" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="40.3229000000000042"/>
             </Node>
            </Node>
           </Node>
           <Node id="67">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="100" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
            </Node>
            <Node id="101">
             <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
             <Node id="128">
              <SimplePredicate field="Age" operator="lessOrEqual" value="30"/>
              <Node id="156" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="22.4646000000000008"/>
              </Node>
              <Node id="157">
               <SimplePredicate field="Fare" operator="greaterThan" value="22.4646000000000008"/>
               <Node id="180" score="0">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
               </Node>
               <Node id="181" score="1">
                <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="129" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="30"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="25">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="42">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;C&quot;</Array>
           </SimpleSetPredicate>
           <Node id="68" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
           </Node>
           <Node id="69" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="37"/>
           </Node>
          </Node>
          <Node id="43">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="70">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="2"/>
            <Node id="102">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="130" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="19"/>
             </Node>
             <Node id="131" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="19"/>
             </Node>
            </Node>
            <Node id="103">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="132" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="12.5"/>
             </Node>
             <Node id="133">
              <SimplePredicate field="Age" operator="greaterThan" value="12.5"/>
              <Node id="158">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
               <Node id="182">
                <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                <Node id="196" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
                </Node>
                <Node id="197">
                 <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
                 <Node id="208">
                  <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;female&quot;</Array>
                  </SimpleSetPredicate>
                  <Node id="218" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="17.875"/>
                  </Node>
                  <Node id="219" score="1">
                   <SimplePredicate field="Fare" operator="greaterThan" value="17.875"/>
                  </Node>
                 </Node>
                 <Node id="209" score="0">
                  <SimpleSetPredicate field="Sex" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;male&quot;</Array>
                  </SimpleSetPredicate>
                 </Node>
                </Node>
               </Node>
               <Node id="183" score="0">
                <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
               </Node>
              </Node>
              <Node id="159">
               <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
               <Node id="184" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
               </Node>
               <Node id="185" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="27"/>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="71" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="2"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="13">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
        <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;Q&quot;</Array>
       </SimpleSetPredicate>
       <Node id="4">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;female&quot;</Array>
        </SimpleSetPredicate>
        <Node id="8">
         <SimplePredicate field="Age" operator="lessOrEqual" value="31.25"/>
         <Node id="16">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="30">
           <SimplePredicate field="Age" operator="lessOrEqual" value="30.25"/>
           <Node id="52">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="78">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="22.7291500000000006"/>
             <Node id="106">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="13.4041499999999996"/>
              <Node id="128" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="20"/>
              </Node>
              <Node id="129">
               <SimplePredicate field="Age" operator="greaterThan" value="20"/>
               <Node id="146" score="1">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
               </Node>
               <Node id="147">
                <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                <Node id="164" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="21.5"/>
                </Node>
                <Node id="165" score="1">
                 <SimplePredicate field="Age" operator="greaterThan" value="21.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
             <Node id="107" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="13.4041499999999996"/>
             </Node>
            </Node>
            <Node id="79" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="22.7291500000000006"/>
            </Node>
           </Node>
           <Node id="53">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            <Node id="80">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="1" type="string">&quot;C&quot;</Array>
             </SimpleSetPredicate>
             <Node id="108" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="14.75"/>
             </Node>
             <Node id="109" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="14.75"/>
             </Node>
            </Node>
            <Node id="81" score="1">
             <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
              <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
             </SimpleSetPredicate>
            </Node>
           </Node>
          </Node>
          <Node id="31" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="30.25"/>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="32">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="1"/>
           <Node id="54">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
            <Node id="82" score="0">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="18.8020500000000013"/>
            </Node>
            <Node id="83" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="18.8020500000000013"/>
            </Node>
           </Node>
           <Node id="55" score="1">
            <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
          <Node id="33" score="1">
           <SimplePredicate field="SibSp" operator="greaterThan" value="1"/>
          </Node>
         </Node>
        </Node>
        <Node id="9" score="1">
         <SimplePredicate field="Age" operator="greaterThan" value="31.25"/>
        </Node>
       </Node>
       <Node id="5">
        <SimpleSetPredicate field="Sex" booleanOperator="isIn">
         <Array n="1" type="string">&quot;male&quot;</Array>
        </SimpleSetPredicate>
        <Node id="10">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="8.48750000000000071"/>
         <Node id="18">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="7.7458499999999999"/>
          <Node id="34">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="56" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="20.5"/>
           </Node>
           <Node id="57">
            <SimplePredicate field="Age" operator="greaterThan" value="20.5"/>
            <Node id="84">
             <SimplePredicate field="Age" operator="lessOrEqual" value="28.25"/>
             <Node id="110">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.47710000000000008"/>
              <Node id="130">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="1" type="string">&quot;C&quot;</Array>
               </SimpleSetPredicate>
               <Node id="148" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
               </Node>
               <Node id="149" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
               </Node>
              </Node>
              <Node id="131" score="0">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
              </Node>
             </Node>
             <Node id="111" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.47710000000000008"/>
             </Node>
            </Node>
            <Node id="85" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="28.25"/>
            </Node>
           </Node>
          </Node>
          <Node id="35" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
          </Node>
         </Node>
         <Node id="19">
          <SimplePredicate field="Fare" operator="greaterThan" value="7.7458499999999999"/>
          <Node id="36" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.78960000000000008"/>
          </Node>
          <Node id="37">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.78960000000000008"/>
           <Node id="58" score="0">
            <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
           </Node>
           <Node id="59">
            <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
            <Node id="86" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
            </Node>
            <Node id="87" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="32"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimplePredicate field="Fare" operator="greaterThan" value="8.48750000000000071"/>
         <Node id="20">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="38">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="28.7250000000000014"/>
           <Node id="60" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="19"/>
           </Node>
           <Node id="61" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="19"/>
           </Node>
          </Node>
          <Node id="39">
           <SimplePredicate field="Fare" operator="greaterThan" value="28.7250000000000014"/>
           <Node id="62">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
            <Node id="88" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
            </Node>
            <Node id="89">
             <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
             <Node id="112" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="113" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
           <Node id="63" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="40">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
           <Node id="64">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="161.19165000000001"/>
            <Node id="90" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="91">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="114" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="28.4125000000000014"/>
             </Node>
             <Node id="115" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="28.4125000000000014"/>
             </Node>
            </Node>
           </Node>
           <Node id="65" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="161.19165000000001"/>
           </Node>
          </Node>
          <Node id="41" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
        <Array n="1" type="string">&quot;S&quot;</Array>
       </SimpleSetPredicate>
       <Node id="6">
        <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
        <Node id="12">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="9.28749999999999964"/>
         <Node id="22" score="0">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;female&quot;</Array>
          </SimpleSetPredicate>
         </Node>
         <Node id="23">
          <SimpleSetPredicate field="Sex" booleanOperator="isIn">
           <Array n="1" type="string">&quot;male&quot;</Array>
          </SimpleSetPredicate>
          <Node id="42">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
           <Node id="66" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="67" score="0">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           </Node>
          </Node>
          <Node id="43" score="0">
           <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimplePredicate field="Fare" operator="greaterThan" value="9.28749999999999964"/>
         <Node id="24" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="8"/>
         </Node>
         <Node id="25">
          <SimplePredicate field="Age" operator="greaterThan" value="8"/>
          <Node id="44">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="68">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="25.0396000000000001"/>
            <Node id="92">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="19.8562499999999993"/>
             <Node id="116">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="10.8166499999999992"/>
              <Node id="132">
               <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
               <Node id="150" score="1">
                <SimplePredicate field="Age" operator="lessOrEqual" value="53.5"/>
               </Node>
               <Node id="151" score="0">
                <SimplePredicate field="Age" operator="greaterThan" value="53.5"/>
               </Node>
              </Node>
              <Node id="133">
               <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
               <Node id="152" score="1">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="9.46875"/>
               </Node>
               <Node id="153" score="0">
                <SimplePredicate field="Fare" operator="greaterThan" value="9.46875"/>
               </Node>
              </Node>
             </Node>
             <Node id="117">
              <SimplePredicate field="Fare" operator="greaterThan" value="10.8166499999999992"/>
              <Node id="134" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
              </Node>
              <Node id="135" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
              </Node>
             </Node>
            </Node>
            <Node id="93" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="19.8562499999999993"/>
            </Node>
           </Node>
           <Node id="69">
            <SimplePredicate field="Fare" operator="greaterThan" value="25.0396000000000001"/>
            <Node id="94" score="1">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="95" score="0">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="45" score="0">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
        <Node id="14">
         <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
         <Node id="26">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
          <Node id="46">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;female&quot;</Array>
           </SimpleSetPredicate>
           <Node id="70">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="23.5"/>
            <Node id="96" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="19.5"/>
            </Node>
            <Node id="97">
             <SimplePredicate field="Age" operator="greaterThan" value="19.5"/>
             <Node id="118">
              <SimplePredicate field="Age" operator="lessOrEqual" value="29.5"/>
              <Node id="136">
               <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
               <Node id="154" score="0">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
               </Node>
               <Node id="155">
                <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                <Node id="166" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="22.5"/>
                </Node>
                <Node id="167" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="22.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="137" score="1">
               <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
              </Node>
             </Node>
             <Node id="119" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="29.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="71" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="23.5"/>
           </Node>
          </Node>
          <Node id="47" score="0">
           <SimpleSetPredicate field="Sex" booleanOperator="isIn">
            <Array n="1" type="string">&quot;male&quot;</Array>
           </SimpleSetPredicate>
          </Node>
         </Node>
         <Node id="27">
          <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
          <Node id="48" score="0">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="10.7979000000000003"/>
          </Node>
          <Node id="49">
           <SimplePredicate field="Fare" operator="greaterThan" value="10.7979000000000003"/>
           <Node id="72">
            <SimplePredicate field="Age" operator="lessOrEqual" value="16"/>
            <Node id="98">
             <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
             <Node id="120" score="1">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="121">
              <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
              <Node id="138" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="1.45999999999999996"/>
              </Node>
              <Node id="139" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="1.45999999999999996"/>
              </Node>
             </Node>
            </Node>
            <Node id="99" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
           <Node id="73">
            <SimplePredicate field="Age" operator="greaterThan" value="16"/>
            <Node id="100">
             <SimplePredicate field="Age" operator="lessOrEqual" value="57"/>
             <Node id="122">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="26.125"/>
              <Node id="140">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="23.2250000000000014"/>
               <Node id="156" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="34"/>
               </Node>
               <Node id="157" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="34"/>
               </Node>
              </Node>
              <Node id="141" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="23.2250000000000014"/>
              </Node>
             </Node>
             <Node id="123">
              <SimplePredicate field="Fare" operator="greaterThan" value="26.125"/>
              <Node id="142">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
               <Node id="158">
                <SimplePredicate field="Age" operator="lessOrEqual" value="32"/>
                <Node id="168">
                 <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
                 <Node id="176" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="28"/>
                 </Node>
                 <Node id="177" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="28"/>
                 </Node>
                </Node>
                <Node id="169">
                 <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
                 <Node id="178" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
                 </Node>
                 <Node id="179" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
                 </Node>
                </Node>
               </Node>
               <Node id="159">
                <SimplePredicate field="Age" operator="greaterThan" value="32"/>
                <Node id="170">
                 <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
                 <Node id="180" score="1">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="48.5"/>
                 </Node>
                 <Node id="181" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="48.5"/>
                 </Node>
                </Node>
                <Node id="171" score="1">
                 <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="143">
               <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
               <Node id="160" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="31.3312499999999972"/>
               </Node>
               <Node id="161">
                <SimplePredicate field="Fare" operator="greaterThan" value="31.3312499999999972"/>
                <Node id="172" score="1">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="39.1437499999999972"/>
                </Node>
                <Node id="173" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="39.1437499999999972"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="101" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="57"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15">
         <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
         <Node id="28">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="2.5"/>
          <Node id="50">
           <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
           <Node id="74" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="75">
            <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
            <Node id="102" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="103">
             <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
             <Node id="124" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
             </Node>
             <Node id="125" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="51">
           <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
           <Node id="76">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;female&quot;</Array>
            </SimpleSetPredicate>
            <Node id="104">
             <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
             <Node id="126">
              <SimplePredicate field="Age" operator="lessOrEqual" value="26"/>
              <Node id="144" score="0">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
              </Node>
              <Node id="145">
               <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
               <Node id="162">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="148.6875"/>
                <Node id="174" score="0">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
                </Node>
                <Node id="175">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
                 <Node id="182" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="19.5999999999999979"/>
                 </Node>
                 <Node id="183" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="19.5999999999999979"/>
                 </Node>
                </Node>
               </Node>
               <Node id="163" score="1">
                <SimplePredicate field="Fare" operator="greaterThan" value="148.6875"/>
               </Node>
              </Node>
             </Node>
             <Node id="127" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="26"/>
             </Node>
            </Node>
            <Node id="105" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
            </Node>
           </Node>
           <Node id="77" score="0">
            <SimpleSetPredicate field="Sex" booleanOperator="isIn">
             <Array n="1" type="string">&quot;male&quot;</Array>
            </SimpleSetPredicate>
           </Node>
          </Node>
         </Node>
         <Node id="29" score="1">
          <SimplePredicate field="Parch" operator="greaterThan" value="2.5"/>
         </Node>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="14">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;female&quot;</Array>
       </SimpleSetPredicate>
       <Node id="4">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;Q&quot;</Array>
        </SimpleSetPredicate>
        <Node id="8" score="0">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="6.98749999999999982"/>
        </Node>
        <Node id="9">
         <SimplePredicate field="Fare" operator="greaterThan" value="6.98749999999999982"/>
         <Node id="16">
          <SimplePredicate field="Parch" operator="lessOrEqual" value="3.5"/>
          <Node id="28">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           <Node id="44">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            <Node id="64" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="46"/>
            </Node>
            <Node id="65">
             <SimplePredicate field="Age" operator="greaterThan" value="46"/>
             <Node id="84" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="55.9354000000000013"/>
             </Node>
             <Node id="85" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="55.9354000000000013"/>
             </Node>
            </Node>
           </Node>
           <Node id="45" score="1">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
          <Node id="29">
           <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
           <Node id="46">
            <SimplePredicate field="Age" operator="lessOrEqual" value="37"/>
            <Node id="66" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="16.5"/>
            </Node>
            <Node id="67">
             <SimplePredicate field="Age" operator="greaterThan" value="16.5"/>
             <Node id="86">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
              <Node id="106">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
               <Node id="120">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                <Node id="128">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;C&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="136" score="1">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="10.84375"/>
                 </Node>
                 <Node id="137" score="0">
                  <SimplePredicate field="Fare" operator="greaterThan" value="10.84375"/>
                 </Node>
                </Node>
                <Node id="129" score="1">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
               </Node>
               <Node id="121" score="1">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
               </Node>
              </Node>
              <Node id="107" score="1">
               <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              </Node>
             </Node>
             <Node id="87" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="47" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="37"/>
           </Node>
          </Node>
         </Node>
         <Node id="17" score="0">
          <SimplePredicate field="Parch" operator="greaterThan" value="3.5"/>
         </Node>
        </Node>
       </Node>
       <Node id="5">
        <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
         <Array n="1" type="string">&quot;S&quot;</Array>
        </SimpleSetPredicate>
        <Node id="10">
         <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
         <Node id="18" score="1">
          <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
         </Node>
         <Node id="19">
          <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
          <Node id="30" score="0">
           <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
          </Node>
          <Node id="31">
           <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
           <Node id="48" score="1">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           </Node>
           <Node id="49">
            <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
            <Node id="68" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="69" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
         <Node id="20">
          <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
          <Node id="32" score="1">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
          </Node>
          <Node id="33">
           <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
           <Node id="50" score="0">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
           </Node>
           <Node id="51" score="1">
            <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="21">
          <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
          <Node id="34">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="17.6000000000000014"/>
           <Node id="52">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="11.4958500000000008"/>
            <Node id="70">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
             <Node id="88">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
              <Node id="108" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
              </Node>
              <Node id="109" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="27"/>
              </Node>
             </Node>
             <Node id="89" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="71" score="1">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
           <Node id="53">
            <SimplePredicate field="Fare" operator="greaterThan" value="11.4958500000000008"/>
            <Node id="72">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="15.1750000000000007"/>
             <Node id="90" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="38.5"/>
             </Node>
             <Node id="91" score="0">
              <SimplePredicate field="Age" operator="greaterThan" value="38.5"/>
             </Node>
            </Node>
            <Node id="73" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="15.1750000000000007"/>
            </Node>
           </Node>
          </Node>
          <Node id="35">
           <SimplePredicate field="Fare" operator="greaterThan" value="17.6000000000000014"/>
           <Node id="54">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="4.5"/>
            <Node id="74">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="20.6625000000000014"/>
             <Node id="92" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="19.125"/>
             </Node>
             <Node id="93" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="19.125"/>
             </Node>
            </Node>
            <Node id="75" score="0">
             <SimplePredicate field="Fare" operator="greaterThan" value="20.6625000000000014"/>
            </Node>
           </Node>
           <Node id="55" score="0">
            <SimplePredicate field="Parch" operator="greaterThan" value="4.5"/>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimpleSetPredicate field="Sex" booleanOperator="isIn">
        <Array n="1" type="string">&quot;male&quot;</Array>
       </SimpleSetPredicate>
       <Node id="6">
        <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
        <Node id="12">
         <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
         <Node id="22" score="1">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
         </Node>
         <Node id="23">
          <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          <Node id="36" score="1">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="20.8249999999999993"/>
          </Node>
          <Node id="37" score="0">
           <SimplePredicate field="Fare" operator="greaterThan" value="20.8249999999999993"/>
          </Node>
         </Node>
        </Node>
        <Node id="13">
         <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
         <Node id="24" score="1">
          <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
         </Node>
         <Node id="25">
          <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
          <Node id="38" score="1">
           <SimplePredicate field="Age" operator="lessOrEqual" value="3.5"/>
          </Node>
          <Node id="39" score="0">
           <SimplePredicate field="Age" operator="greaterThan" value="3.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7">
        <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
        <Node id="14">
         <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
         <Node id="26">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="1" type="string">&quot;C&quot;</Array>
          </SimpleSetPredicate>
          <Node id="40">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           <Node id="56">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="369.927099999999996"/>
            <Node id="76">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="94">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="13.3687500000000004"/>
              <Node id="110" score="1">
               <SimplePredicate field="Age" operator="lessOrEqual" value="25"/>
              </Node>
              <Node id="111" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="25"/>
              </Node>
             </Node>
             <Node id="95">
              <SimplePredicate field="Fare" operator="greaterThan" value="13.3687500000000004"/>
              <Node id="112" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="77.9646000000000043"/>
              </Node>
              <Node id="113" score="0">
               <SimplePredicate field="Fare" operator="greaterThan" value="77.9646000000000043"/>
              </Node>
             </Node>
            </Node>
            <Node id="77" score="0">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
           <Node id="57" score="1">
            <SimplePredicate field="Fare" operator="greaterThan" value="369.927099999999996"/>
           </Node>
          </Node>
          <Node id="41">
           <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           <Node id="58" score="1">
            <SimplePredicate field="Age" operator="lessOrEqual" value="13.5"/>
           </Node>
           <Node id="59">
            <SimplePredicate field="Age" operator="greaterThan" value="13.5"/>
            <Node id="78">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="49.2541999999999973"/>
             <Node id="96">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="19.8708500000000008"/>
              <Node id="114" score="0">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="14.8499999999999996"/>
              </Node>
              <Node id="115" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="14.8499999999999996"/>
              </Node>
             </Node>
             <Node id="97" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="19.8708500000000008"/>
             </Node>
            </Node>
            <Node id="79">
             <SimplePredicate field="Fare" operator="greaterThan" value="49.2541999999999973"/>
             <Node id="98" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="99.9895999999999958"/>
             </Node>
             <Node id="99" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="99.9895999999999958"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="27">
          <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
           <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
          </SimpleSetPredicate>
          <Node id="42">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           <Node id="60" score="0">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="26.1437500000000007"/>
           </Node>
           <Node id="61">
            <SimplePredicate field="Fare" operator="greaterThan" value="26.1437500000000007"/>
            <Node id="80" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="30.75"/>
            </Node>
            <Node id="81">
             <SimplePredicate field="Fare" operator="greaterThan" value="30.75"/>
             <Node id="100" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="51.2479000000000013"/>
             </Node>
             <Node id="101">
              <SimplePredicate field="Fare" operator="greaterThan" value="51.2479000000000013"/>
              <Node id="116">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="3"/>
               <Node id="122" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;Q&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="123">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
                <Node id="130">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="86.7374999999999972"/>
                 <Node id="138" score="0">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="49"/>
                 </Node>
                 <Node id="139" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="49"/>
                 </Node>
                </Node>
                <Node id="131" score="1">
                 <SimplePredicate field="Fare" operator="greaterThan" value="86.7374999999999972"/>
                </Node>
               </Node>
              </Node>
              <Node id="117" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="3"/>
              </Node>
             </Node>
            </Node>
           </Node>
          </Node>
          <Node id="43">
           <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           <Node id="62">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="47.7479000000000013"/>
            <Node id="82">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
             <Node id="102">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
              <Node id="118">
               <SimplePredicate field="Age" operator="lessOrEqual" value="61.5"/>
               <Node id="124">
                <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                <Node id="132" score="0">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="1" type="string">&quot;Q&quot;</Array>
                 </SimpleSetPredicate>
                </Node>
                <Node id="133">
                 <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                  <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                 </SimpleSetPredicate>
                 <Node id="140">
                  <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
                  <Node id="146">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="11"/>
                   <Node id="156" score="0">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="5.25"/>
                   </Node>
                   <Node id="157" score="0">
                    <SimplePredicate field="Fare" operator="greaterThan" value="5.25"/>
                   </Node>
                  </Node>
                  <Node id="147" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="11"/>
                  </Node>
                 </Node>
                 <Node id="141">
                  <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                  <Node id="148">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="7.19585000000000008"/>
                   <Node id="158" score="0">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="23"/>
                   </Node>
                   <Node id="159" score="1">
                    <SimplePredicate field="Age" operator="greaterThan" value="23"/>
                   </Node>
                  </Node>
                  <Node id="149" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="7.19585000000000008"/>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="125">
                <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                <Node id="134">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="32.5"/>
                 <Node id="142">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="8.77500000000000036"/>
                  <Node id="150" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="7.86459999999999937"/>
                  </Node>
                  <Node id="151" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="7.86459999999999937"/>
                  </Node>
                 </Node>
                 <Node id="143">
                  <SimplePredicate field="Fare" operator="greaterThan" value="8.77500000000000036"/>
                  <Node id="152">
                   <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
                   <Node id="160" score="0">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
                   </Node>
                   <Node id="161" score="1">
                    <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
                   </Node>
                  </Node>
                  <Node id="153">
                   <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
                   <Node id="162" score="1">
                    <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                   </Node>
                   <Node id="163" score="0">
                    <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="135">
                 <SimplePredicate field="Age" operator="greaterThan" value="32.5"/>
                 <Node id="144" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="7.98750000000000071"/>
                 </Node>
                 <Node id="145">
                  <SimplePredicate field="Fare" operator="greaterThan" value="7.98750000000000071"/>
                  <Node id="154" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="8.35624999999999929"/>
                  </Node>
                  <Node id="155">
                   <SimplePredicate field="Fare" operator="greaterThan" value="8.35624999999999929"/>
                   <Node id="164">
                    <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                    <Node id="166" score="0">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="40.5"/>
                    </Node>
                    <Node id="167" score="0">
                     <SimplePredicate field="Age" operator="greaterThan" value="40.5"/>
                    </Node>
                   </Node>
                   <Node id="165" score="0">
                    <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
               </Node>
              </Node>
              <Node id="119">
               <SimplePredicate field="Age" operator="greaterThan" value="61.5"/>
               <Node id="126" score="1">
                <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
               </Node>
               <Node id="127" score="0">
                <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
               </Node>
              </Node>
             </Node>
             <Node id="103" score="0">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
            <Node id="83">
             <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             <Node id="104" score="0">
              <SimplePredicate field="Age" operator="lessOrEqual" value="27"/>
             </Node>
             <Node id="105" score="1">
              <SimplePredicate field="Age" operator="greaterThan" value="27"/>
             </Node>
            </Node>
           </Node>
           <Node id="63" score="0">
            <SimplePredicate field="Fare" operator="greaterThan" value="47.7479000000000013"/>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="15" score="0">
         <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
        </Node>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
   <Segment id="15">
    <True/>
    <TreeModel modelName="randomForest_Model" functionName="classification" algorithmName="randomForest" splitCharacteristic="binarySplit">
     <MiningSchema>
      <MiningField name="Survived" usageType="predicted"/>
      <MiningField name="PassengerId" usageType="active"/>
      <MiningField name="Pclass" usageType="active"/>
      <MiningField name="Name" usageType="active"/>
      <MiningField name="Sex" usageType="active"/>
      <MiningField name="Age" usageType="active"/>
      <MiningField name="SibSp" usageType="active"/>
      <MiningField name="Parch" usageType="active"/>
      <MiningField name="Ticket" usageType="active"/>
      <MiningField name="Fare" usageType="active"/>
      <MiningField name="Cabin" usageType="active"/>
      <MiningField name="Embarked" usageType="active"/>
     </MiningSchema>
     <Node id="1">
      <True/>
      <Node id="2">
       <SimplePredicate field="Age" operator="lessOrEqual" value="5.5"/>
       <Node id="4" score="1">
        <SimplePredicate field="Fare" operator="lessOrEqual" value="20.8249999999999993"/>
       </Node>
       <Node id="5">
        <SimplePredicate field="Fare" operator="greaterThan" value="20.8249999999999993"/>
        <Node id="8" score="0">
         <SimplePredicate field="Fare" operator="lessOrEqual" value="21.5499999999999972"/>
        </Node>
        <Node id="9">
         <SimplePredicate field="Fare" operator="greaterThan" value="21.5499999999999972"/>
         <Node id="12">
          <SimplePredicate field="Age" operator="lessOrEqual" value="2.5"/>
          <Node id="18">
           <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
           <Node id="30" score="0">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;Q&quot;</Array>
            </SimpleSetPredicate>
           </Node>
           <Node id="31">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="48" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="3"/>
            </Node>
            <Node id="49" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="3"/>
            </Node>
           </Node>
          </Node>
          <Node id="19">
           <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
           <Node id="32" score="1">
            <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
           </Node>
           <Node id="33" score="0">
            <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
           </Node>
          </Node>
         </Node>
         <Node id="13">
          <SimplePredicate field="Age" operator="greaterThan" value="2.5"/>
          <Node id="20" score="1">
           <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
          </Node>
          <Node id="21" score="0">
           <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
          </Node>
         </Node>
        </Node>
       </Node>
      </Node>
      <Node id="3">
       <SimplePredicate field="Age" operator="greaterThan" value="5.5"/>
       <Node id="6">
        <SimplePredicate field="SibSp" operator="lessOrEqual" value="3.5"/>
        <Node id="10">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;female&quot;</Array>
         </SimpleSetPredicate>
         <Node id="14">
          <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
          <Node id="22">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="26.125"/>
           <Node id="34">
            <SimplePredicate field="Fare" operator="lessOrEqual" value="20.25"/>
            <Node id="50" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="12.8249999999999993"/>
            </Node>
            <Node id="51">
             <SimplePredicate field="Fare" operator="greaterThan" value="12.8249999999999993"/>
             <Node id="72" score="1">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="73" score="1">
              <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             </Node>
            </Node>
           </Node>
           <Node id="35">
            <SimplePredicate field="Fare" operator="greaterThan" value="20.25"/>
            <Node id="52">
             <SimplePredicate field="Age" operator="lessOrEqual" value="40"/>
             <Node id="74" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="22"/>
             </Node>
             <Node id="75" score="1">
              <SimplePredicate field="Fare" operator="greaterThan" value="22"/>
             </Node>
            </Node>
            <Node id="53" score="0">
             <SimplePredicate field="Age" operator="greaterThan" value="40"/>
            </Node>
           </Node>
          </Node>
          <Node id="23" score="1">
           <SimplePredicate field="Fare" operator="greaterThan" value="26.125"/>
          </Node>
         </Node>
         <Node id="15">
          <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
          <Node id="24">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="1" type="string">&quot;Q&quot;</Array>
           </SimpleSetPredicate>
           <Node id="36">
            <SimplePredicate field="Age" operator="lessOrEqual" value="29.25"/>
            <Node id="54" score="1">
             <SimplePredicate field="Age" operator="lessOrEqual" value="17"/>
            </Node>
            <Node id="55">
             <SimplePredicate field="Age" operator="greaterThan" value="17"/>
             <Node id="76" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.68125000000000036"/>
             </Node>
             <Node id="77">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.68125000000000036"/>
              <Node id="92" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="7.78960000000000008"/>
              </Node>
              <Node id="93" score="1">
               <SimplePredicate field="Fare" operator="greaterThan" value="7.78960000000000008"/>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="37" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="29.25"/>
           </Node>
          </Node>
          <Node id="25">
           <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
            <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
           </SimpleSetPredicate>
           <Node id="38">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="1" type="string">&quot;C&quot;</Array>
            </SimpleSetPredicate>
            <Node id="56">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="17.252049999999997"/>
             <Node id="78" score="0">
              <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             </Node>
             <Node id="79">
              <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
              <Node id="94" score="0">
               <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
              </Node>
              <Node id="95" score="0">
               <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
              </Node>
             </Node>
            </Node>
            <Node id="57" score="1">
             <SimplePredicate field="Fare" operator="greaterThan" value="17.252049999999997"/>
            </Node>
           </Node>
           <Node id="39">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="58" score="1">
             <SimplePredicate field="Fare" operator="lessOrEqual" value="7.70000000000000018"/>
            </Node>
            <Node id="59">
             <SimplePredicate field="Fare" operator="greaterThan" value="7.70000000000000018"/>
             <Node id="80" score="0">
              <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
             </Node>
             <Node id="81" score="0">
              <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
             </Node>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
        <Node id="11">
         <SimpleSetPredicate field="Sex" booleanOperator="isIn">
          <Array n="1" type="string">&quot;male&quot;</Array>
         </SimpleSetPredicate>
         <Node id="16">
          <SimplePredicate field="Fare" operator="lessOrEqual" value="26.2687500000000007"/>
          <Node id="26">
           <SimplePredicate field="Fare" operator="lessOrEqual" value="7.9104000000000001"/>
           <Node id="40">
            <SimplePredicate field="Age" operator="lessOrEqual" value="32.5"/>
            <Node id="60" score="0">
             <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="61">
             <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
             <Node id="82" score="0">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="2.00625000000000009"/>
             </Node>
             <Node id="83">
              <SimplePredicate field="Fare" operator="greaterThan" value="2.00625000000000009"/>
              <Node id="96">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
               </SimpleSetPredicate>
               <Node id="104">
                <SimplePredicate field="Age" operator="lessOrEqual" value="28.75"/>
                <Node id="116">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="20.5"/>
                 <Node id="128" score="0">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="6.98959999999999937"/>
                 </Node>
                 <Node id="129" score="1">
                  <SimplePredicate field="Fare" operator="greaterThan" value="6.98959999999999937"/>
                 </Node>
                </Node>
                <Node id="117">
                 <SimplePredicate field="Age" operator="greaterThan" value="20.5"/>
                 <Node id="130">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="1" type="string">&quot;C&quot;</Array>
                  </SimpleSetPredicate>
                  <Node id="138" score="0">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="7.22710000000000008"/>
                  </Node>
                  <Node id="139" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="7.22710000000000008"/>
                  </Node>
                 </Node>
                 <Node id="131">
                  <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                   <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                  </SimpleSetPredicate>
                  <Node id="140">
                   <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                   <Node id="150" score="0">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="24.5"/>
                   </Node>
                   <Node id="151">
                    <SimplePredicate field="Age" operator="greaterThan" value="24.5"/>
                    <Node id="162" score="0">
                     <SimplePredicate field="Fare" operator="lessOrEqual" value="7.74375000000000036"/>
                    </Node>
                    <Node id="163" score="0">
                     <SimplePredicate field="Fare" operator="greaterThan" value="7.74375000000000036"/>
                    </Node>
                   </Node>
                  </Node>
                  <Node id="141" score="0">
                   <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                  </Node>
                 </Node>
                </Node>
               </Node>
               <Node id="105" score="1">
                <SimplePredicate field="Age" operator="greaterThan" value="28.75"/>
               </Node>
              </Node>
              <Node id="97">
               <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
               </SimpleSetPredicate>
               <Node id="106" score="0">
                <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
               </Node>
               <Node id="107">
                <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                <Node id="118" score="0">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="23.5"/>
                </Node>
                <Node id="119" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="23.5"/>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
           </Node>
           <Node id="41" score="0">
            <SimplePredicate field="Age" operator="greaterThan" value="32.5"/>
           </Node>
          </Node>
          <Node id="27">
           <SimplePredicate field="Fare" operator="greaterThan" value="7.9104000000000001"/>
           <Node id="42">
            <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
            <Node id="62">
             <SimplePredicate field="Parch" operator="lessOrEqual" value="0.5"/>
             <Node id="84" score="1">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="7.98750000000000071"/>
             </Node>
             <Node id="85">
              <SimplePredicate field="Fare" operator="greaterThan" value="7.98750000000000071"/>
              <Node id="98">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="13.25"/>
               <Node id="108">
                <SimplePredicate field="Age" operator="lessOrEqual" value="30.5"/>
                <Node id="120">
                 <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                 <Node id="132">
                  <SimplePredicate field="Fare" operator="lessOrEqual" value="9.83539999999999992"/>
                  <Node id="142">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
                   <Node id="152">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="8.83125000000000071"/>
                    <Node id="164">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="26.5"/>
                     <Node id="170" score="0">
                      <SimplePredicate field="Fare" operator="lessOrEqual" value="8.17500000000000071"/>
                     </Node>
                     <Node id="171" score="0">
                      <SimplePredicate field="Fare" operator="greaterThan" value="8.17500000000000071"/>
                     </Node>
                    </Node>
                    <Node id="165" score="1">
                     <SimplePredicate field="Age" operator="greaterThan" value="26.5"/>
                    </Node>
                   </Node>
                   <Node id="153" score="0">
                    <SimplePredicate field="Fare" operator="greaterThan" value="8.83125000000000071"/>
                   </Node>
                  </Node>
                  <Node id="143">
                   <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
                   <Node id="154" score="0">
                    <SimplePredicate field="Fare" operator="lessOrEqual" value="9.49164999999999992"/>
                   </Node>
                   <Node id="155">
                    <SimplePredicate field="Fare" operator="greaterThan" value="9.49164999999999992"/>
                    <Node id="166" score="0">
                     <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                    </Node>
                    <Node id="167" score="1">
                     <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                    </Node>
                   </Node>
                  </Node>
                 </Node>
                 <Node id="133">
                  <SimplePredicate field="Fare" operator="greaterThan" value="9.83539999999999992"/>
                  <Node id="144" score="0">
                   <SimplePredicate field="Age" operator="lessOrEqual" value="27.5"/>
                  </Node>
                  <Node id="145">
                   <SimplePredicate field="Age" operator="greaterThan" value="27.5"/>
                   <Node id="156" score="0">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="28.5"/>
                   </Node>
                   <Node id="157" score="0">
                    <SimplePredicate field="Age" operator="greaterThan" value="28.5"/>
                   </Node>
                  </Node>
                 </Node>
                </Node>
                <Node id="121" score="0">
                 <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                </Node>
               </Node>
               <Node id="109">
                <SimplePredicate field="Age" operator="greaterThan" value="30.5"/>
                <Node id="122">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="12.7624999999999993"/>
                 <Node id="134">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="44.5"/>
                  <Node id="146">
                   <SimplePredicate field="Fare" operator="lessOrEqual" value="8.20625000000000071"/>
                   <Node id="158" score="1">
                    <SimplePredicate field="Age" operator="lessOrEqual" value="33"/>
                   </Node>
                   <Node id="159" score="0">
                    <SimplePredicate field="Age" operator="greaterThan" value="33"/>
                   </Node>
                  </Node>
                  <Node id="147" score="0">
                   <SimplePredicate field="Fare" operator="greaterThan" value="8.20625000000000071"/>
                  </Node>
                 </Node>
                 <Node id="135" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="44.5"/>
                 </Node>
                </Node>
                <Node id="123" score="0">
                 <SimplePredicate field="Fare" operator="greaterThan" value="12.7624999999999993"/>
                </Node>
               </Node>
              </Node>
              <Node id="99">
               <SimplePredicate field="Fare" operator="greaterThan" value="13.25"/>
               <Node id="110" score="0">
                <SimplePredicate field="Fare" operator="lessOrEqual" value="18.3937500000000007"/>
               </Node>
               <Node id="111">
                <SimplePredicate field="Fare" operator="greaterThan" value="18.3937500000000007"/>
                <Node id="124" score="0">
                 <SimplePredicate field="Fare" operator="lessOrEqual" value="19.3770999999999987"/>
                </Node>
                <Node id="125">
                 <SimplePredicate field="Fare" operator="greaterThan" value="19.3770999999999987"/>
                 <Node id="136">
                  <SimplePredicate field="Age" operator="lessOrEqual" value="33"/>
                  <Node id="148" score="1">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="1" type="string">&quot;Q&quot;</Array>
                   </SimpleSetPredicate>
                  </Node>
                  <Node id="149">
                   <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                    <Array n="3" type="string">&quot;&quot;   &quot;C&quot;   &quot;S&quot;</Array>
                   </SimpleSetPredicate>
                   <Node id="160" score="0">
                    <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
                   </Node>
                   <Node id="161">
                    <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
                    <Node id="168" score="0">
                     <SimplePredicate field="SibSp" operator="lessOrEqual" value="1.5"/>
                    </Node>
                    <Node id="169" score="0">
                     <SimplePredicate field="SibSp" operator="greaterThan" value="1.5"/>
                    </Node>
                   </Node>
                  </Node>
                 </Node>
                 <Node id="137" score="0">
                  <SimplePredicate field="Age" operator="greaterThan" value="33"/>
                 </Node>
                </Node>
               </Node>
              </Node>
             </Node>
            </Node>
            <Node id="63">
             <SimplePredicate field="Parch" operator="greaterThan" value="0.5"/>
             <Node id="86" score="1">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="1" type="string">&quot;C&quot;</Array>
              </SimpleSetPredicate>
             </Node>
             <Node id="87" score="0">
              <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
               <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
              </SimpleSetPredicate>
             </Node>
            </Node>
           </Node>
           <Node id="43">
            <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
            <Node id="64" score="1">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
            </Node>
            <Node id="65" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
            </Node>
           </Node>
          </Node>
         </Node>
         <Node id="17">
          <SimplePredicate field="Fare" operator="greaterThan" value="26.2687500000000007"/>
          <Node id="28">
           <SimplePredicate field="Pclass" operator="lessOrEqual" value="1.5"/>
           <Node id="44">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;C&quot;   &quot;Q&quot;</Array>
            </SimpleSetPredicate>
            <Node id="66">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="0.5"/>
             <Node id="88" score="1">
              <SimplePredicate field="Age" operator="lessOrEqual" value="19.5"/>
             </Node>
             <Node id="89">
              <SimplePredicate field="Age" operator="greaterThan" value="19.5"/>
              <Node id="100">
               <SimplePredicate field="Parch" operator="lessOrEqual" value="1.5"/>
               <Node id="112" score="0">
                <SimplePredicate field="Age" operator="lessOrEqual" value="31"/>
               </Node>
               <Node id="113">
                <SimplePredicate field="Age" operator="greaterThan" value="31"/>
                <Node id="126" score="1">
                 <SimplePredicate field="Age" operator="lessOrEqual" value="36.5"/>
                </Node>
                <Node id="127" score="0">
                 <SimplePredicate field="Age" operator="greaterThan" value="36.5"/>
                </Node>
               </Node>
              </Node>
              <Node id="101" score="0">
               <SimplePredicate field="Parch" operator="greaterThan" value="1.5"/>
              </Node>
             </Node>
            </Node>
            <Node id="67">
             <SimplePredicate field="SibSp" operator="greaterThan" value="0.5"/>
             <Node id="90">
              <SimplePredicate field="Fare" operator="lessOrEqual" value="98.7520999999999987"/>
              <Node id="102" score="1">
               <SimplePredicate field="Fare" operator="lessOrEqual" value="89.5520999999999958"/>
              </Node>
              <Node id="103">
               <SimplePredicate field="Fare" operator="greaterThan" value="89.5520999999999958"/>
               <Node id="114" score="1">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="1" type="string">&quot;C&quot;</Array>
                </SimpleSetPredicate>
               </Node>
               <Node id="115" score="0">
                <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
                 <Array n="3" type="string">&quot;&quot;   &quot;Q&quot;   &quot;S&quot;</Array>
                </SimpleSetPredicate>
               </Node>
              </Node>
             </Node>
             <Node id="91" score="0">
              <SimplePredicate field="Fare" operator="greaterThan" value="98.7520999999999987"/>
             </Node>
            </Node>
           </Node>
           <Node id="45">
            <SimpleSetPredicate field="Embarked" booleanOperator="isIn">
             <Array n="2" type="string">&quot;&quot;   &quot;S&quot;</Array>
            </SimpleSetPredicate>
            <Node id="68" score="0">
             <SimplePredicate field="SibSp" operator="lessOrEqual" value="2.5"/>
            </Node>
            <Node id="69" score="0">
             <SimplePredicate field="SibSp" operator="greaterThan" value="2.5"/>
            </Node>
           </Node>
          </Node>
          <Node id="29">
           <SimplePredicate field="Pclass" operator="greaterThan" value="1.5"/>
           <Node id="46" score="0">
            <SimplePredicate field="Pclass" operator="lessOrEqual" value="2.5"/>
           </Node>
           <Node id="47">
            <SimplePredicate field="Pclass" operator="greaterThan" value="2.5"/>
            <Node id="70" score="0">
             <SimplePredicate field="Age" operator="lessOrEqual" value="21"/>
            </Node>
            <Node id="71" score="1">
             <SimplePredicate field="Age" operator="greaterThan" value="21"/>
            </Node>
           </Node>
          </Node>
         </Node>
        </Node>
       </Node>
       <Node id="7" score="0">
        <SimplePredicate field="SibSp" operator="greaterThan" value="3.5"/>
       </Node>
      </Node>
     </Node>
    </TreeModel>
   </Segment>
  </Segmentation>
 </MiningModel>
</PMML>
`

func TestRandomForest(t *testing.T) {
	randomForest := []byte(randomForestXml)
	var rf goscore.RandomForest
	err := xml.Unmarshal(randomForest, &rf)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(rf.Trees))
	fmt.Println(rf.Trees[0].Attrs)

	//for _, tt := range TreeTests {
	//	actual := goscore.TraverseTree(n, tt.features)
	//	if actual != tt.score {
	//		t.Errorf("expected %f, actual %f",
	//			tt.score,
	//			actual)
	//	}
	//}
}
